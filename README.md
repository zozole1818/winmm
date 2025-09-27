# winmm => Windows Mouse Mover

It's a simple Go program that moves the mouse every 60 seconds. First 200px right, then left.
Additionally, after mouse gets back to initial position - left windows key is pressed twice. This closes the cycle and start a new one.

You can set different time schedule but simply providing a number when running the command.

## Installation
```bash
$ go install github.com/zozole1818/winmm@latest
```

## Usage
```bash
$ winmm # will trigger mouse movement every 60 seconds
$ winmm 5 # will trigger mouse movement every 5 seconds
```

Note: `winmm` will only work on Windows. Under the hood Go uses win32 API.