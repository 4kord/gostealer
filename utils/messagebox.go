package utils

import (
	"syscall"
	"unsafe"
)

func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	captioni, _ := syscall.UTF16PtrFromString(caption)
	titlei, _ := syscall.UTF16PtrFromString(title)

	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(captioni)),
		uintptr(unsafe.Pointer(titlei)),
		uintptr(flags))

	return int(ret)
}

func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK|16)
}
