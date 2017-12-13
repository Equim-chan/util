package cli

import (
	"os"
)

// ClearTerminal clears the terminal screen and move cursor to top-left corner.
//
// This function assumes stdout is a tty but it does not check it. You must
// check that (by calling either github.com/mattn/go-isatty.IsTerminal or
// golang.org/x/crypto/ssh/terminal.IsTerminal on os.Stdout) on your own before
// actually calling ClearTerminal.
//
// Unlike other common simple solutions that just calls "cls.exe" or "clear(1)",
// it calls Console API (Windows) or writes VT100 escape characters (POSIX)
// directly.
func ClearTerminal() error {
	return clearTerminal()
}

// NewTTYReader allocates an extra tty that can be used to read.
//
// NewTTYReader and NewTTYWriter are useful for prompting when stdin/stdout is
// redirected. Since it returns a *os.File, you can also call
// golang.org/x/crypto/ssh/terminal.ReadPassword on it! An example:
/*
	package main

	import (
		"crypto/subtle"
		"fmt"
		"io"
		"log"
		"os"

		"ekyu.moe/util/cli"
		"golang.org/x/crypto/ssh/terminal"
	)

	func main() {
		r, err := cli.NewTTYReader()
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		w, err := cli.NewTTYWriter()
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()

		// isatty check
		fmt.Fprintln(
			w,
			terminal.IsTerminal(int(os.Stdin.Fd())),
			terminal.IsTerminal(int(os.Stdout.Fd())),
		)
		fmt.Fprintln(
			w,
			terminal.IsTerminal(int(r.Fd())),
			terminal.IsTerminal(int(w.Fd())),
		)

		// Basic prompt
		fmt.Fprint(w, "Would you like to continue? [y/N] ")
		var ans string
		if fmt.Fscanln(r, &ans); ans != "y" {
			fmt.Fprintln(w, "You chose to abort.")
			os.Exit(1)
		}
		fmt.Fprintln(w, "You chose to continue.")

		// Simple passphrase prompt
		fmt.Fprint(w, "Enter the passphrase: ")
		passphrase, err := terminal.ReadPassword(int(r.Fd()))
		fmt.Fprintln(w)
		if err != nil {
			log.Fatal(err)
		}

		// Of course in real usage you must do hash/kdf
		if subtle.ConstantTimeCompare(passphrase, []byte("a secret")) == 0 {
			fmt.Fprintln(w, "Wrong passphrase.")
			os.Exit(1)
		}

		// Do something like encrypt/decrypt or encode/decode through stdin and
		// stdout.
		io.Copy(os.Stdout, os.Stdin)
	}
*/
// Run it:
//     $ go run main.go <<< "raw data blablabla" | cat
//     false false
//     true true
//     Would you like to continue? [y/N] y
//     You chose to continue.
//     Enter the passphrase:
//     raw data blablabla
func NewTTYReader() (*os.File, error) {
	return newTTYReader()
}

// NewTTYReader allocates an extra tty that can be used to write.
func NewTTYWriter() (*os.File, error) {
	return newTTYWriter()
}
