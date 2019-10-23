package discovery

import (
	"io"
	"log"
	"os/user"

	"golang.org/x/sys/windows/registry"
)

// CurrentUser returns the current user
func CurrentUser() ([]map[string]string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	a := make(map[string]string)
	a["Username: "] = u.Username
	a["SID: "] = u.Uid
	a["Home Directory: "] = u.HomeDir

	userInfo := make([]map[string]string, 0)
	userInfo = append(userInfo, a)

	return userInfo, err
}

// ActiveUsers returns
func ActiveUsers() []*user.User {
	// query registry key, "HKEY_USERS", to get list of security IDs (SIDs)
	sids := queryRegistrySIDs()

	// parse list of SIDs, looking for those that are 46 characters long, indicating active user accounts
	activeUserSids := parseSIDs(sids)

	// lookup user information for each SID
	activeAccounts := lookupUserAccount(activeUserSids)

	return activeAccounts

}

func queryRegistrySIDs() []string {
	// open registry key, "HKEY_USERS"
	k, err := registry.OpenKey(registry.USERS, "", registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	// Read HKEY_USERS subkeys to get user SIDs
	userSIDs, err := k.ReadSubKeyNames(1024)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	return userSIDs
}

func parseSIDs(sids []string) []string {
	activeUserSIDs := make([]string, 0)
	for _, sid := range sids {
		if len(sid) == 46 {
			// append SID
			activeUserSIDs = append(activeUserSIDs, sid)
		}
	}
	return activeUserSIDs
}

func lookupUserAccount(activeUserSids []string) []*user.User {
	activeUsers := make([]*user.User, 0)
	for _, s := range activeUserSids {
		userInfo, _ := user.LookupId(s)
		activeUsers = append(activeUsers, userInfo)
	}
	return activeUsers
}
