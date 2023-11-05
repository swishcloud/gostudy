package messagehook

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	WH_KEYBOARD_LL = 13
)

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")

	user32, _            = syscall.LoadLibrary("user32.dll")
	setWindowsHookExA, _ = syscall.GetProcAddress(user32, "SetWindowsHookExW")
	getMessageW, _       = syscall.GetProcAddress(user32, "GetMessageW")
)

type (
	DWORD     uint32
	WPARAM    uintptr
	LPARAM    uintptr
	LRESULT   uintptr
	HANDLE    uintptr
	HINSTANCE HANDLE
	HHOOK     HANDLE
	HWND      HANDLE
	UINT      uint32
	ULONG_PTR uintptr
)
type CWPRETSTRUCT struct {
	lResult LRESULT
	lParam  LPARAM
	wParam  WPARAM
	message UINT
	hwnd    HWND
}

type POINT struct {
	X, Y int32
}
type Msg struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}
type LPMSG *Msg
type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

func SetWindowsHookExA(idHook int, lpfn HOOKPROC, hmod HINSTANCE, dwThreadId DWORD) (r HHOOK) {
	var nargs uintptr = 4
	l := syscall.NewCallback(lpfn)
	if r1, _, err := syscall.Syscall6(uintptr(setWindowsHookExA), nargs, WH_KEYBOARD_LL, l, uintptr(hmod), 0, 0, 0); err != 0 {
		abort("Call SetWindowsHookExA", err)
	} else {
		r = HHOOK(r1)
	}
	return
}

func GetMessageW(lpMsg LPMSG, hWnd HWND, wMsgFilterMin UINT, wMsgFilterMax UINT) (r int) {
	var nargs uintptr = 4
	if r1, _, err := syscall.Syscall6(uintptr(getMessageW), nargs, uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd), uintptr(wMsgFilterMax), uintptr(wMsgFilterMax), 0, 0); err != 0 {
		abort("Call GetMessageW", err)
	} else {
		r = int(r1)
	}
	return
}

func GetModuleHandle() (handle uintptr) {
	var nargs uintptr = 0
	if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
		abort("Call GetModuleHandle", callErr)
	} else {
		handle = ret
	}
	return
}
func Start(h HOOKPROC) {
	r := SetWindowsHookExA(WH_KEYBOARD_LL, h, HINSTANCE(GetModuleHandle()), 0)
	if r == 0 {
		panic("unexpected return value from SetWindowsHookExA:" + string(r))
	}
	var msg Msg
	GetMessageW(LPMSG(&msg), 0, 0, 0)
}
