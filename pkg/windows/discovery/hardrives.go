package discovery

import (
	"github.com/shirou/gopsutil/disk"
)

// GetDrives iterates through the alphabet to return a list of mounted drives
func GetDrives() ([]disk.PartitionStat, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	return partitions, err
}
