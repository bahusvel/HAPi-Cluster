package main

import (
	"os/exec"
	"sync"

	"github.com/bahusvel/ClusterPipe/common"
)

var thisCPD = common.CPD{}

var procMutex = sync.RWMutex{}
var processes = map[string][]common.PreparedTask{}

func Run(task common.PreparedTask) error {
	cmd := exec.Command(task.Command, task.Args...)
	var err error
	if task.Stderr != nil {
		cmd.Stderr, err = task.Stderr.Open()
		if err != nil {
			return err
		}
	}
	if task.Stdout != nil {
		cmd.Stdout, err = task.Stdout.Open()
		if err != nil {
			return err
		}
	}
	if task.Stdin != nil {
		cmd.Stdin, err = task.Stdin.Open()
		if err != nil {
			return err
		}
	}

	procMutex.Lock()
	defer procMutex.Unlock()

	err = cmd.Start()
	if err != nil {
		return err
	}
	task.Process = cmd
	if jobProcs, ok := processes[task.Job]; !ok {
		processes[task.Job] = []common.PreparedTask{task}
	} else {
		processes[task.Job] = append(jobProcs, task)
	}

	return nil
}

func main() {
	go common.RunPipeServer()
	Start()
}
