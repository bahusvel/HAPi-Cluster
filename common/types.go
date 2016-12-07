package common

import (
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

const (
	CP_ROOT  = "/var/clusterpipe/"
	CP_PIPES = CP_ROOT + "pipes/"
)

type ResourceQuota struct {
	CPU  int
	RAM  int
	DISK int
}

/*
type SchedulingRequirements struct {
}
*/

type DataItem struct {
}

type Task struct {
	Command string
	Args    []string
	Stdin   string
	Stdout  string
	Stderr  string
	Quota   ResourceQuota
	//SchedReq SchedulingRequirements
}

type PreparedTask struct {
	Job     string
	Command string
	Args    []string
	Stdin   *FIFO
	Stdout  *FIFO
	Stderr  *FIFO
	Process *exec.Cmd
}

type Job struct {
	Owner string
	Name  string
	Quota ResourceQuota
	Flow  []Task
}

type CPDStatus struct {
	LastCheckin time.Time
	Jobs        []string
	CPUUtil     float64
	CPUTime     []cpu.TimesStat
	MemStat     mem.VirtualMemoryStat
	NetStat     net.IOCountersStat
}

type CPD struct {
	Host          string
	HostInfo      *host.InfoStat
	CPUInfo       []cpu.InfoStat
	NetInfo       []net.InterfaceStat
	CurrentStatus *CPDStatus
}
