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

// NewTTY allocates an extra tty that can be used to read or write.
//
// This is useful for prompting when stdin/stdout is a pipe, for example:
//     $ some-encoding-tool < bin.dat > bin.asc
//     Please enter your passphrase:
//     Comfirm your passphrase:
//     Done, 16K -> 20K.
func NewTTY() (*os.File, error) {
	return newTTY()
}
