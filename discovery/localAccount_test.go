package discovery

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	cmd := exec.Command("whoami")
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Error running cmd.CombinedOutput: %q", err)
	}
	currentUser := strings.TrimSuffix(string(cmdOutput), "\n")

	want := string(currentUser)
	got, err := GetCurrentUser()
	if err != nil {
		t.Errorf("Error running GetCurrentUser: %q", err)
	}
	if got != want {
		t.Errorf("GetCurrentUser(): got %q, expected %q", got, want)
	}
}
