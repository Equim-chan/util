package cli

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
