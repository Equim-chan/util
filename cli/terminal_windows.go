package cli

import (
	"os"
	"syscall"
	"unsafe"

	. "ekyu.moe/util/cli/internal"
	. "ekyu.moe/util/sugar"
)

// References:
//     https://www.zybuluo.com/Gestapo/note/32082#程序-33gotoxy与-clrscr函数
//     https://support.microsoft.com/en-us/help/99261/how-to-performing-clear-screen-cls-in-a-console-application
func clearTerminal() error {
	console := syscall.Stdout
	coordScreen := Coord{0, 0}

	// get the number of character cells in the current buffer
	csbi, err := GetConsoleScreenBufferInfo(console)
	if err != nil {
		return err
	}
	conSize := uint32(csbi.Size.X * csbi.Size.Y)

	// fill the entire screen with blanks
	if _, err := FillConsoleOutputCharacter(console, ' ', conSize, coordScreen); err != nil {
		return err
	}

	// get the current text attribute
	if csbi, err = GetConsoleScreenBufferInfo(console); err != nil {
		return err
	}

	// now set the buffer's attributes accordingly
	if _, err := FillConsoleOutputAttribute(console, csbi.Attributes, conSize, coordScreen); err != nil {
		return err
	}

	// put the cursor at (0, 0)
	if err := SetConsoleCursorPosition(console, coordScreen); err != nil {
		return err
	}

	return nil
}

func newTTYReader() (*os.File, error) {
	sa := syscall.SecurityAttributes{}
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 0x01 // TRUE

	handle, err := syscall.CreateFile(
		Must2(syscall.UTF16PtrFromString("CONIN$")).(*uint16),
		syscall.GENERIC_READ,
		syscall.FILE_SHARE_READ,
		&sa,
		syscall.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(handle), "/dev/tty"), nil
}

func newTTYWriter() (*os.File, error) {
	sa := syscall.SecurityAttributes{}
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 0x01 // TRUE

	handle, err := syscall.CreateFile(
		Must2(syscall.UTF16PtrFromString("CONOUT$")).(*uint16),
		syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_WRITE,
		&sa,
		syscall.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	return os.NewFile(uintptr(handle), "/dev/tty"), nil
}
