package winmm

import "golang.org/x/sys/windows"

var (
	user32       = windows.NewLazySystemDLL("user32.dll")
	setCursorPos = user32.NewProc("SetCursorPos")
	getCursorPos = user32.NewProc("GetCursorPos")
	sendInput    = user32.NewProc("SendInput")
)
