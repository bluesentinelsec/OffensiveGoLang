package discovery

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

// Procs returns a slice of process objects
// Look at "PrintProcSummary()" and "PrintProcDetails()" for examples on displaying the process info
func Procs() ([][]string, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}
	processes := make([][]string, 0)
	for _, ps := range procs {
		// assign each process value to a variable
		i := ps.Pid
		pid := strconv.Itoa(int(i))

		i, _ = ps.Ppid()
		pPid := strconv.Itoa(int(i))

		pName, _ := ps.Name()
		pCmdLine, _ := ps.Cmdline()

		//pCreate, _ := ps.CreateTime()

		pUser, err := ps.Username()
		if err != nil {
			pUser = "Access Denied"
		}

		p := make([]string, 0)
		p = append(p, pid, pPid, pName, pCmdLine, pUser)

		// append each process, p, to the processes slice
		processes = append(processes, p)
	}
	return processes, nil
}

// PrintProcSummary provides summarized process information in a table
func PrintProcSummary(procs [][]string) {
	// display table headers
	fmt.Printf("| %-8s | %-8s \t | %-32s \t| %-16s\n", "[PID]", "[Parent]", "[Name]", "[User]")

	// display process info in columns
	for i := range procs {
		fmt.Printf("| %-8s | %-8s \t | %-32s \t| %-16s\n", procs[i][0], procs[i][1], procs[i][2], procs[i][4])
	}

}

// PrintProcDetails provides detailed process information in a list
func PrintProcDetails(procs [][]string) {
	for i := range procs {
		fmt.Printf("Name:            %s\n", procs[i][2])
		fmt.Printf("User:            %s\n", procs[i][4])
		fmt.Printf("PID:             %s\n", procs[i][0])
		fmt.Printf("Parent:          %s\n", procs[i][1])
		fmt.Printf("Command Line:    %s\n\n", procs[i][3])
	}
}

// PrintProcessTree displays each process in a tree
func PrintProcessTree() {
	fmt.Println("Sorry, this function is incomplete.")
	/*
		for each proc
			if proc is in a blacklist:
				continue
			else:
				print proc info

			if proc has children:
				for each child process:
					print child process info
					add children to blacklist
	*/
}
