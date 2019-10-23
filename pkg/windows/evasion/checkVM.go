package evasion

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows/registry"
)

//==========================================================
//	VM Process Checks
//==========================================================

// CheckVMprocs searches for VM associated processes on the target host
func CheckVMprocs() ([]string, error) {
	// get process list
	procs, err := listProcesses()
	if err != nil {
		return nil, err
	}

	// parse process list for VM-associated processes
	vmProcs := parseProcesses(procs)
	return vmProcs, err
}

// listProcesses returns process names in a list
func listProcesses() ([]string, error) {
	// list processes
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
	// store each process name in a list
	procNames := make([]string, 0)
	for _, proc := range procs {
		s, _ := proc.Name()
		procNames = append(procNames, s)
	}
	return procNames, err
}

//  parseProcesses returns virtualization processes
func parseProcesses(procs []string) []string {
	// list of processes known to be run in a VM
	badProcs := []string{
		"vmacthlp.exe",
		"vmtoolsd.exe"}

	// return running VM processes in a list, vmProcs
	vmProcs := make([]string, 0)
	for _, proc := range procs {
		for _, badProc := range badProcs {
			if proc == badProc {
				vmProcs = append(vmProcs, badProc)
			}
		}
	}
	return vmProcs
}

//==========================================================
//	VM Filesystem Artifact Checks
//==========================================================

// CheckVMfiles searches for VM associated files on the target host
func CheckVMfiles() []string {
	vmFiles := []string{
		"C:\\Program Files\\VMware",
		"C:\\Program Files (x86)\\VMware"}
	s := make([]string, 0)
	for _, file := range vmFiles {
		_, err := os.Stat(file)
		if err != nil {
			continue
		}
		s = append(s, file)
	}
	if s != nil {
		return s
	}
	return nil
}

func displayVMfiles(s []string) {
	fmt.Println("[!] Virtual machine files found:")
	for _, f := range s {
		fmt.Println("[!]", f)
	}
}

//==========================================================
//	VM Registry Checks
//==========================================================

// CheckVMregistry queries the Windows registry for VM-associated entries
func CheckVMregistry() (string, error) {
	path := `HARDWARE\DESCRIPTION\System\BIOS`
	name := "SystemProductName"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue(name)
	if err != nil {
		return "", err
	}
	if s == "VMware, Inc." {
		return s, err
	}
	return "", err

}

//==========================================================
//	VM Service Checks
//==========================================================

// CheckVMservices queries for VM-associated services
func CheckVMservices() {
	/*
	   Services:
	   Running  VMTools            VMware Tools
	   Running  VGAuthService      VMware Alias Manager and Ticket Ser
	   Stopped  vmvss              VMware Snapshot Provider
	   Running  VMware Physical... VMware Physical Disk Helper Service
	   Stopped  VMwareCAFCommAm... VMware CAF AMQP Communication Service
	   Stopped  VMwareCAFManage... VMware CAF Management Agent Service


	   Get-PnpDevice | Select-String VMWare
	*/
}
