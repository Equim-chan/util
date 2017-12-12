package cli

import (
	"os"
	"syscall"
	"unsafe"
)

// References:
//     https://www.zybuluo.com/Gestapo/note/32082#程序-33gotoxy与-clrscr函数
//     https://support.microsoft.com/en-us/help/99261/how-to-performing-clear-screen-cls-in-a-console-application
func clearTerminal() error {
	console := syscall.Stdout
	coordScreen := coord{0, 0}

	// get the number of character cells in the current buffer
	csbi, err := getConsoleScreenBufferInfo(console)
	if err != nil {
		return err
	}
	conSize := uint32(csbi.Size.X * csbi.Size.Y)

	// fill the entire screen with blanks
	if _, err := fillConsoleOutputCharacter(console, ' ', conSize, coordScreen); err != nil {
		return err
	}

	// get the current text attribute
	if csbi, err = getConsoleScreenBufferInfo(console); err != nil {
		return err
	}

	// now set the buffer's attributes accordingly
	if _, err := fillConsoleOutputAttribute(console, csbi.Attributes, conSize, coordScreen); err != nil {
		return err
	}

	// put the cursor at (0, 0)
	if err := setConsoleCursorPosition(console, coordScreen); err != nil {
		return err
	}

	return nil
}

func newTTY() (*os.File, error) {
	sa := syscall.SecurityAttributes{}
	sa.Length = uint32(unsafe.Sizeof(sa))
	sa.InheritHandle = 0x01 // TRUE

	// won't return err because the string given is constrant and doesn't contain a NUL
	conin, _ := syscall.UTF16PtrFromString("CONIN$")

	handle, err := syscall.CreateFile(
		conin,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
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

type (
	/*
		WORD   uint16
		DWORD  uint32
		TCHAR  rune
	*/

	coord struct {
		X, Y int16
	}

	smallRect struct {
		Left, Top, Right, Bottom int16
	}

	consoleScreenBufferInfo struct {
		Size              coord
		CursorPosition    coord
		Attributes        uint16
		Window            smallRect
		MaximumWindowSize coord
	}
)

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procFillConsoleOutputAttribute = kernel32.NewProc("FillConsoleOutputAttribute")
	procFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
	procSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
)

// Little endian
func coordToUint32(c coord) uint32 {
	return uint32(int32(c.Y)<<16 + int32(c.X))
}

func getConsoleScreenBufferInfo(console syscall.Handle) (csbi *consoleScreenBufferInfo, err error) {
	csbi = new(consoleScreenBufferInfo)

	r, _, e := procGetConsoleScreenBufferInfo.Call(
		uintptr(console),
		uintptr(unsafe.Pointer(csbi)),
	)
	if r == 0 {
		err = e
	}

	return
}

func fillConsoleOutputAttribute(console syscall.Handle, attribute uint16, length uint32, writeCoord coord) (numberOfAttrsWritten uint32, err error) {
	r, _, e := procFillConsoleOutputAttribute.Call(
		uintptr(console),
		uintptr(attribute),
		uintptr(length),
		uintptr(coordToUint32(writeCoord)),
		uintptr(unsafe.Pointer(&numberOfAttrsWritten)),
	)
	if r == 0 {
		err = e
	}

	return
}

func fillConsoleOutputCharacter(console syscall.Handle, character rune, length uint32, writeCoord coord) (numberOfAttrsWritten uint32, err error) {
	r, _, e := procFillConsoleOutputCharacter.Call(
		uintptr(console),
		uintptr(character),
		uintptr(length),
		uintptr(coordToUint32(writeCoord)),
		uintptr(unsafe.Pointer(&numberOfAttrsWritten)),
	)
	if r == 0 {
		err = e
	}

	return
}

func setConsoleCursorPosition(console syscall.Handle, cursorPosition coord) (err error) {
	r, _, e := procSetConsoleCursorPosition.Call(
		uintptr(console),
		uintptr(coordToUint32(cursorPosition)),
	)
	if r == 0 {
		err = e
	}

	return
}
