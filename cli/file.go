package cli

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

var (
	ErrAbortedByUser = errors.New("aborted by user")
)

// Access a file (not a directory), return the file info and error.
// Filename "-" is considered as stdin, and AccessFile will return (nil, nil) in this case.
func AccessFile(filename string) (os.FileInfo, error) {
	if filename == "-" {
		return nil, nil
	}

	stat, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return stat, fmt.Errorf(`File "%s" does not exist`, filename)
		}
		return stat, err
	}

	if stat.IsDir() {
		return stat, fmt.Errorf(`"%s" is a directory, not a file`, filename)
	}

	return stat, nil
}

type Input struct {
	*os.File
	isStdin    bool
	isBuffered bool
}

// Same as AccessFile, but also open the file.
// If filename is "-", read from stdin, and file info will be nil in this case.
func AccessOpenFile(filename string) (*Input, os.FileInfo, error) {
	if filename == "-" {
		return &Input{os.Stdin, true, false}, nil, nil
	}

	stat, err := AccessFile(filename)
	if err != nil {
		return nil, nil, err
	}

	file, err := os.Open(filename)
	return &Input{file, false, false}, stat, err
}

// Same as AccessOpenFile, but buffers stdin to a temp file when filename is "-", and
// returns the *Input and os.FileInfo based on the temp file. It will resume the stdin
// after reading an EOF from it, instead of closing stdin.
func AccessOpenFileBuffered(filename string) (*Input, os.FileInfo, error) {
	if filename != "-" {
		return AccessOpenFile(filename)
	}

	file, err := ioutil.TempFile("", "stdin")
	if err != nil {
		return nil, nil, err
	}

	if _, err := io.Copy(file, os.Stdin); err != nil {
		os.Remove(file.Name())
		return nil, nil, err
	}

	// resume os.Stdin after read all
	os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")

	// seek back to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		return nil, nil, err
	}

	// stat
	stat, err := file.Stat()
	if err != nil {
		os.Remove(file.Name())
		return nil, nil, err
	}

	return &Input{file, true, true}, stat, nil
}

func (i *Input) IsStdin() bool {
	return i.isStdin
}

// Close the underlying *os.File.
// If *os.File is os.Stdin, remove the underlying temp file.
func (bs *Input) Close() error {
	if bs.isStdin {
		if !bs.isBuffered {
			return nil
		}
		if err := bs.File.Close(); err != nil {
			os.Remove(bs.Name())
			return err
		}
		return os.Remove(bs.Name())
	}

	return bs.File.Close()
}

// Check if a file already exists, if true, prompt for overriding. On err == nil,
// either there is no filename conflict or the user permitted overriding, or
// the filename is "-", which is considered as stdout.
func PromptOverride(filename string) error {
	if filename == "-" {
		return nil
	}

	stat, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if stat.IsDir() {
		fmt.Printf(`File "%s" already exists, override? (y/N) `, filename)
	} else {
		fmt.Printf(`"%s" already exists and it is a directory! Override? (y/N) `, filename)
	}
	r := ""
	if fmt.Scanln(&r); r != "y" {
		return ErrAbortedByUser
	}

	return nil
}

type Output struct {
	*os.File
	isStdout bool
}

// Same as PromptOverride, but also open the file with flag and perm passed to os.OpenFile.
// When filename is "-", write to stdout.
func PromptOverrideOpen(filename string, flag int, perm os.FileMode) (*Output, error) {
	if filename == "-" {
		return &Output{os.Stdout, true}, nil
	}

	if err := PromptOverride(filename); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return &Output{file, false}, nil
}

// Short cut for PromptOverrideOpen with os.Create.
func PromptOverrideCreate(filename string) (*Output, error) {
	return PromptOverrideOpen(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func (o *Output) IsStdout() bool {
	return o.isStdout
}

// Close the underlying *os.File.
// If the *os.File is os.Stdout, it will not actually close the stdout.
func (o *Output) Close() error {
	if !o.isStdout {
		return o.File.Close()
	}

	// does not actually close stdout
	return nil
}

// Parse a set of globs to file name list, considering "-" as stdin/stdout.
// The only possible returned error is ErrBadPattern for path/filepath.Glob().
// When an error is encountered, the currently parsed filenames will be returned.
// If ignoreInvalid is true, error will always be nil.
func ParseFileList(args []string, ignoreInvalid bool) ([]string, error) {
	if len(args) == 0 {
		return []string{"-"}, nil
	}

	filelist := []string{}
	for _, arg := range args {
		if arg == "-" {
			filelist = append(filelist, "-")
			continue
		}

		matches, err := filepath.Glob(arg)
		if err != nil {
			if !ignoreInvalid {
				return filelist, err
			}
			continue
		}

		filelist = append(filelist, matches...)
	}

	return filelist, nil
}
