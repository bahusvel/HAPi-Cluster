package common

import (
	"log"
	"testing"

	"github.com/shirou/gopsutil/disk"
)

func TestDiskStatThingy(t *testing.T) {
	counterMap, err := disk.IOCounters()
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(counterMap)
}
