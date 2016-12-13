package common

import (
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

const (
	CP_ROOT  = "/var/clusterpipe/"
	CP_PIPES = CP_ROOT + "pipes/"
)

type TaskID int64

func (id TaskID) IsValid() bool {
	return id >= 0
}

type Task struct {
	TID     TaskID
	Node    string
	Command string
	Args    []string
	Stdin   *FIFO
	Stdout  *FIFO
	Stderr  *FIFO
	Process *exec.Cmd
}

type CPDStatus struct {
	LastCheckin time.Time
	Jobs        []string
	CPUUtil     float64
	CPUTime     []cpu.TimesStat
	MemStat     mem.VirtualMemoryStat
	NetStat     net.IOCountersStat
	DiskUsage   []disk.UsageStat
	DiskStat    []disk.IOCountersStat
}

type CPD struct {
	Host          string
	HostInfo      *host.InfoStat
	CPUInfo       []cpu.InfoStat
	NetInfo       []net.InterfaceStat
	CurrentStatus *CPDStatus
}
