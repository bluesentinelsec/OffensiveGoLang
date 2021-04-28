package discovery

import (
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	// get current user through a shell command
	cmd := exec.Command("whoami")
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error running cmd.CombinedOutput: %q", err)
	}
	// trim carriage return and new line characters
	var currentUser string
	if runtime.GOOS == "windows" {
		currentUser = strings.TrimSuffix(string(cmdOutput), "\r\n")
	} else {
		currentUser = strings.TrimSuffix(string(cmdOutput), "\n")
	}

	// run the test
	want := string(currentUser)
	got, err := GetCurrentUser()
	if err != nil {
		t.Errorf("Error running GetCurrentUser: %q", err)
	}
	if got != want {
		t.Errorf("GetCurrentUser(): got %q, expected %q", got, want)
	}
}
