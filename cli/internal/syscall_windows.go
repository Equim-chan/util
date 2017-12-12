// Extra syscalls for Windows
package internal

import (
	"syscall"
	"unsafe"
)

type (
	Handle = syscall.Handle
	/*
		WORD   uint16
		DWORD  uint32
		TCHAR  rune
	*/

	Coord struct {
		X, Y int16
	}

	SmallRect struct {
		Left, Top, Right, Bottom int16
	}

	ConsoleScreenBufferInfo struct {
		Size              Coord
		CursorPosition    Coord
		Attributes        uint16
		Window            SmallRect
		MaximumWindowSize Coord
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
func coordToUint32(c Coord) uint32 {
	return uint32(int32(c.Y)<<16 + int32(c.X))
}

func GetConsoleScreenBufferInfo(console Handle) (csbi *ConsoleScreenBufferInfo, err error) {
	csbi = new(ConsoleScreenBufferInfo)

	r, _, e := procGetConsoleScreenBufferInfo.Call(
		uintptr(console),
		uintptr(unsafe.Pointer(csbi)),
	)
	if r == 0 {
		err = e
	}

	return
}

func FillConsoleOutputAttribute(console Handle, attribute uint16, length uint32, writeCoord Coord) (numberOfAttrsWritten uint32, err error) {
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

func FillConsoleOutputCharacter(console Handle, character rune, length uint32, writeCoord Coord) (numberOfAttrsWritten uint32, err error) {
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

func SetConsoleCursorPosition(console Handle, cursorPosition Coord) (err error) {
	r, _, e := procSetConsoleCursorPosition.Call(
		uintptr(console),
		uintptr(coordToUint32(cursorPosition)),
	)
	if r == 0 {
		err = e
	}

	return
}
