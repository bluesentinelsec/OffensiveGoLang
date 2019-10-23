package execution

import (
	"os/exec"
	"syscall"
)

// RunCMD executes a command with cmd.exe
func RunCMD(cmd string) ([]byte, error) {
	c := exec.Command("cmd.exe", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := c.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return output, err

}
