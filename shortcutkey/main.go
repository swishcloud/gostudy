package main

import (
	"fmt"
	"unsafe"

	"github.com/swishcloud/gostudy/messagehook"
)

type (
	DWORD     uint32
	ULONG_PTR uint32
)

const (
	WM_KEYUP      = 0x0101
	WM_KEYDOWN    = 0x0100
	WM_SYSKEYUP   = 0x105
	WM_SYSKEYDOWN = 0x104
)

type KBDLLHOOKSTRUCT struct {
	vkCode      DWORD
	scanCode    DWORD
	flags       DWORD
	time        DWORD
	dwExtraInfo ULONG_PTR
}

var keys = []int{}

func main() {
	messagehook.Start(func(i int, wparam messagehook.WPARAM, lparam messagehook.LPARAM) messagehook.LRESULT {
		data := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
		if wparam == WM_KEYDOWN || wparam == WM_SYSKEYDOWN {
			keys = append(keys, int(data.vkCode))
			if len(keys) > 3 {
				keys = []int{}
			}
			OnKeyDown()
		} else if wparam == WM_KEYUP || wparam == WM_SYSKEYUP {
			for i, v := range keys {
				if v == int(data.vkCode) {
					keys = append(keys[:i], keys[i+1:]...)
					break
				}
			}
		}
		return messagehook.LRESULT(0)
	})
}

func OnKeyDown() {
	val := 0
	for _, v := range keys {
		val = val | v
	}
	fmt.Println(keys, val)

	//shit+alt+a shortcut
	if val == 229 {
		//todo something
		fmt.Println("todo something")
	}
}
