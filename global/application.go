package global

import "runtime"

// get the currently running system type only once
var runTime = runtime.GOOS

const (
	LINUX   = "linux"
	OSX     = "darwin"
	WINDOWS = "windows"
)

func GetRuntime() string {
	return runTime
}

func IsLinux() bool {
	return runTime == LINUX
}

func IsMac() bool {
	return runTime == OSX
}

func IsWindows() bool {
	return runTime == WINDOWS
}
