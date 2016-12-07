package main

import (
	"log"
	"time"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

const (
	STAT_INTERVAL = 5 * time.Second
)

func GatherHostInfo() error {
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

func StartStatMonitor(callback func(common.CPDStatus)) error {
	newStatus := common.CPDStatus{}
	thisCPD.CurrentStatus = &newStatus
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
		}
		newStatus.MemStat = *tmpMem
		netA, err := net.IOCounters(false)
		if err != nil {
			log.Println("Failed to obtain network statistics", err)
		}
		newStatus.NetStat = netA[0]

		callback(newStatus)
		time.Sleep(STAT_INTERVAL)
	}
}
