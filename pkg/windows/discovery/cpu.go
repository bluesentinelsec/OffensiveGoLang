package discovery

import (
	"golang.org/x/sys/windows/registry"
)

// GetCPUmodel returns the target system's CPU model
func GetCPUmodel() (string, error) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "Identifier"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	v, _, err := k.GetStringValue(regValue)
	if err != nil {
		return "", err
	}
	return v, err
}

// GetCPUname returns the target system's CPU name
func GetCPUname() (string, error) {
	key := `HARDWARE\DESCRIPTION\System\CentralProcessor\0`
	regValue := "ProcessorNameString"

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	v, _, err := k.GetStringValue(regValue)
	if err != nil {
		return "", err
	}
	return v, err
}
