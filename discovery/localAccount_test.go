package discovery

import "testing"

func TestGetCurrentUser(t *testing.T) {
	want := "administrator"
	got := GetCurrentUser()
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
