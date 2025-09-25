package winmm

import (
	"fmt"
	"unsafe"
)

type Point struct {
	X, Y int32
}

func GetMousePosition() (Point, error) {
	var pt Point
	r1, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if r1 == 0 {
		return pt, fmt.Errorf("getCursorPos: %v", err)
	}
	return pt, nil
}

func MoveMouse(pt Point) error {
	r1, _, err := setCursorPos.Call(uintptr(pt.X), uintptr(pt.Y))
	if r1 == 0 {
		return fmt.Errorf("SetCursorPos: %v", err)
	}
	return nil
}
