package evasion

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

// CheckDCOtoolProcs checks the target system for defensive-tool processes
func CheckDCOtoolProcs() ([]string, error) {
	// make a list of known-bad processes
	badDCOprocs := []string{
		"procdump.exe",
		"procdump64.exe",
		"procexp.exe",
		"procexp64.exe",
		"Procmon.exe",
		"Sysmon.exe",
		"Sysmon64.exe",
		"Tcpview.exe",
		"Autoruns.exe",
		"Autoruns64.exe"}

	// get process list
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

	// compare known-bad to running process list
	badProcs := make([]string, 0)
	for _, badDCOproc := range badDCOprocs {
		for _, proc := range procNames {
			if badDCOproc == proc {
				badProcs = append(badProcs, badDCOproc)
			}
		}
	}
	return badProcs, err
}

// CheckDCOtoolFiles searches the file system for defensive tools
func CheckDCOtoolFiles() ([]string, error) {
	// make a list of known-bad DCO files
	badDCOfiles := []string{
		"SysinternalsSuite",
		"Autoruns.exe"}
	// walk the file system
	// compare known-bad files to file system
	fmt.Println(badDCOfiles)
	return nil, nil
}
