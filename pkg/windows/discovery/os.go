package discovery

import (
	"golang.org/x/sys/windows/registry"
)

// OSinfo provides basic information about the target operating system
type OSinfo struct {
	ProductName      string
	ReleaseID        string
	CurrentBuild     string
	InstallationType string
	RegisteredOwner  string
	InstallDate      uint64
	InstallTime      uint64
	ProductID        string
}

// GetOSinfo returns information about the target system OS
func GetOSinfo() (OSinfo, error) {
	var osInfo OSinfo
	key := `SOFTWARE\Microsoft\Windows NT\CurrentVersion`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return osInfo, err
	}
	defer k.Close()
	osInfo.InstallationType, _, _ = k.GetStringValue("InstallationType")
	osInfo.ProductID, _, _ = k.GetStringValue("ProductId")
	osInfo.ProductName, _, _ = k.GetStringValue("ProductName")
	osInfo.RegisteredOwner, _, _ = k.GetStringValue("RegisteredOwner")
	osInfo.ReleaseID, _, _ = k.GetStringValue("ReleaseId")
	osInfo.CurrentBuild, _, _ = k.GetStringValue("CurrentBuild")
	// ToDo: convert epoc times to readable format
	osInfo.InstallDate, _, _ = k.GetIntegerValue("InstallDate")
	osInfo.InstallTime, _, _ = k.GetIntegerValue("InstallTime")

	return osInfo, nil
}

/*
Other registry keys of interest:
HKEY_LOCAL_MACHINE\HARDWARE\DESCRIPTION\System\BIOS
	SystemManufacturer
		ASUSTeK COMPUTER INC.
	SystemProductName
		FX503VM

HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Control\Session Manager\Environment
	PROCESSOR_ARCHITECTURE
		amd64
HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Control\TimeZoneInformation
	TimeZoneKeyName
		Eastern Standard Time

USB Info?
	HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\USB
	HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Enum\USBSTOR
*/
