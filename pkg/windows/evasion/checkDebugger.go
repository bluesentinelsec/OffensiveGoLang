package evasion

import (
	"syscall"
)

var (
	kernel32DLL       = syscall.NewLazyDLL("kernel32.dll")
	isDebuggerPresent = kernel32DLL.NewProc("IsDebuggerPresent")
)

// CheckDebugger determines if the current process is being examined in a debugger
func CheckDebugger() (syscall.Handle, error) {
	r1, _, err := isDebuggerPresent.Call()
	if err != nil {
		return syscall.Handle(r1), err
	}

	// returns '1' if process is in a debugger
	return syscall.Handle(r1), nil

}
