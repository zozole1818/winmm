package main

import (
	"fmt"
	"time"
	"unsafe"
)

// Key flags
const (
	inputKeyboard  = 1
	keyeventfKeyup = 0x0002
	keyLwin        = 0x5B
)

type input struct {
	Type    uint32
	Ki      keybdinput // we only use keyboard input
	Padding uint32     // padding for 64-bit alignment
}

type keybdinput struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

func PressLeftWindows() error {
	input := input{
		Type: inputKeyboard,
		Ki: keybdinput{
			WVk: keyLwin,
		},
	}
	// Key down
	err := sendKeyboardInput(input)
	if err != nil {
		return fmt.Errorf("left windows press failed: %v", err)
	}
	<-time.After(5 * time.Millisecond)

	// Key up
	input.Ki.DwFlags = keyeventfKeyup
	err = sendKeyboardInput(input)
	if err != nil {
		return fmt.Errorf("left windows release failed: %v", err)
	}

	return nil
}

func sendKeyboardInput(input input) error {
	r1, _, err := sendInput.Call(
		1,
		uintptr(unsafe.Pointer(&input)),
		unsafe.Sizeof(input),
	)
	if r1 == 0 {
		return fmt.Errorf("SendInput failed: %v", err)
	}
	return nil
}
