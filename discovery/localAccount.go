package discovery

import "os/user"

// GetCurrentUser returns the program's current user account
// https://attack.mitre.org/techniques/T1087/001/
func GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	userName := currentUser.Username
	return userName, err
}
