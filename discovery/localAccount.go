// This package provides functions for Account Discovery: Local Account TTPs
// https://attack.mitre.org/techniques/T1087/001/
package discovery

import "C"

import "os/user"

// GetCurrentUser returns the program's current user account.
func GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	userName := currentUser.Username
	return userName, err
}

// GetCurrentUserID returns the current user's ID on Linux-based systems.
// On Windows systems, this function returns a security identifier (SID).
func GetCurrentUserID() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	userID := currentUser.Uid
	return userID, err
}

// GetCurrentGroupID gets the current user's primary group ID.
// On Windows systems, this function returns a security identifier (SID).
func GetCurrentGroupID() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	groupID := currentUser.Gid
	return groupID, err
}

// GetUserHomeDir returns the path to the current user's home directory
func GetUserHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := currentUser.HomeDir
	return homeDir, err
}
