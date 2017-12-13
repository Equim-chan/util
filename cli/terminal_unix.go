// +build !windows

package cli

import (
	"os"
)

func clearTerminal() error {
	_, err := os.Stdout.Write([]byte{0x1b, '[', '2', 'J', 0x1b, '[', 'H'})

	return err
}

func newTTYReader() (*os.File, error) {
	return os.OpenFile("/dev/tty", os.O_RDONLY, 0)
}

func newTTYWriter() (*os.File, error) {
	return os.OpenFile("/dev/tty", os.O_WRONLY, 0)
}
