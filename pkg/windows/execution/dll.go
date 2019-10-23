package execution

import (
	"syscall"
)

// WinDLLProc executes a windows DLL function
func WinDLLProc(dllName string, funcName string, args ...uintptr) (uintptr, error) {
	mod := syscall.NewLazyDLL(dllName)
	proc := mod.NewProc(funcName)
	ret, _, err := proc.Call(args...)
	return ret, err
}

/*
Example using this function:

	ret, err := execution.WinDLLProc("user32.dll",
		"MessageBoxW",
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
		uintptr(0x00000003))
	fmt.Println(ret, err)

*/
