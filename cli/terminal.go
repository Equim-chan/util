package cli

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear the terminal screen.
// On windows, "cls" is called, while on others "clear" is called.
func ClearScreen() error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	return cmd.Run()
}
