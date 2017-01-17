package common

import (
	"io"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
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
	TID            TaskID
	Node           string
	Scheduler      string
	SchedulingArgs []interface{}
	Command        string
	Args           []string
}

type ScheduleTask struct {
	Task
	Stdio   TaskPipes
	Process *exec.Cmd
}

type CPDStatus struct {
	LastCheckin time.Time
	Jobs        []string
	CPUUtil     float64
	LoadStat    *load.AvgStat
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

type TaskPipes struct {
	Stderr io.ReadWriteCloser
	Stdout io.ReadWriteCloser
	Stdin  io.ReadWriteCloser
}
