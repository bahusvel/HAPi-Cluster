package common

import (
	"log"
	"sort"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

const (
	STAT_INTERVAL = 5 * time.Second
)

var MonitorDiskUsage = false
var partitions []disk.PartitionStat

func GatherHostInfo(thisCPD *CPD) error {
	var err error
	thisCPD.HostInfo, err = host.Info()
	if err != nil {
		return err
	}
	thisCPD.CPUInfo, err = cpu.Info()
	if err != nil {
		return err
	}
	thisCPD.NetInfo, err = net.Interfaces()
	if err != nil {
		return err
	}
	return nil
}

func getDiskStatus(newStatus *CPDStatus) {
	newStatus.DiskUsage = []disk.UsageStat{}
	for _, partition := range partitions {
		tmpDisk, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Println("Failed to obtain disk usage stats", err)
		} else {
			newStatus.DiskUsage = append(newStatus.DiskUsage, *tmpDisk)
		}
	}

	ioCounters, err := disk.IOCounters()
	if err != nil {
		log.Println("Failed to obtain disk IO counters", err)
		return
	}

	devices := []string{}
	for device, _ := range ioCounters {
		devices = append(devices, device)
	}
	sort.Strings(devices)
	newStatus.DiskStat = []disk.IOCountersStat{}
	for _, device := range devices {
		newStatus.DiskStat = append(newStatus.DiskStat, ioCounters[device])
	}
}

func StartStatMonitor(callback func(CPDStatus)) error {
	newStatus := CPDStatus{}
	var err error
	partitions, err = disk.Partitions(false)
	if err != nil {
		return err
	}
	for {
		cpuA, err := cpu.Percent(0, false)
		if err != nil {
			log.Println("Failed to obtain CPU load", err)
		}
		newStatus.CPUUtil = cpuA[0]
		newStatus.CPUTime, err = cpu.Times(true)
		if err != nil {
			log.Println("Failed to obtain CPU load", err)
		}

		tmpMem, err := mem.VirtualMemory()
		if err != nil {
			log.Println("Failed to obtain memory statistics", err)
		} else {
			newStatus.MemStat = *tmpMem
		}

		netA, err := net.IOCounters(false)
		if err != nil {
			log.Println("Failed to obtain network statistics", err)
		} else {
			newStatus.NetStat = netA[0]
		}

		if MonitorDiskUsage {
			getDiskStatus(&newStatus)
		}

		callback(newStatus)
		time.Sleep(STAT_INTERVAL)
	}
}
