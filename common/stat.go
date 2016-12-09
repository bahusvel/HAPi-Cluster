package common

import (
	"log"
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

func StartStatMonitor(callback func(CPDStatus)) error {
	newStatus := CPDStatus{}
	partitions, err := disk.Partitions(false)
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

			newStatus.DiskUsage = []disk.UsageStat{}
			for _, partition := range partitions {
				var tmpDisk *disk.UsageStat
				tmpDisk, err = disk.Usage(partition.Mountpoint)
				if err != nil {
					log.Println("Failed to obtain disk usage stats", err)
				} else {
					newStatus.DiskUsage = append(newStatus.DiskUsage, *tmpDisk)
				}
			}

			ioCounters, err := disk.IOCounters()
			if err != nil {
				log.Println("Failed to obtain disk IO counters", err)
			} else {
				newStatus.DiskStat = []disk.IOCountersStat{}
				for _, counter := range ioCounters {
					newStatus.DiskStat = append(newStatus.DiskStat, counter)
				}
			}

		}

		callback(newStatus)
		time.Sleep(STAT_INTERVAL)
	}
}
