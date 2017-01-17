package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/bahusvel/ClusterPipe/kissrpc"
)

const (
	COM_PORT = ":3335"
)

func init() {
	kissrpc.RegisterType(common.CPD{})
	kissrpc.RegisterType(common.CPDStatus{})
}

func prepareTask(task common.Task) error {
	if !task.TID.IsValid() {
		return fmt.Errorf("Invalid task id")
	}
	cmd := exec.Command(task.Command, task.Args...)
	procMutex.Lock()
	defer procMutex.Unlock()
	scheduledTask := common.ScheduleTask{Task: task, Process: cmd}
	processes[task.TID] = &scheduledTask
	return nil
}

func startTask(taskID common.TaskID) error {
	if !taskID.IsValid() {
		return fmt.Errorf("Invalid task id")
	}
	task, ok := processes[taskID]
	if !ok {
		return fmt.Errorf("Invalid task id")
	}
	task.Process.Stderr = task.Stdio.Stderr
	task.Process.Stdin = task.Stdio.Stdin
	task.Process.Stdout = task.Stdio.Stdout

	err := task.Process.Start()
	if err != nil {
		return err
	}
	return nil
}

func taskKill(id common.TaskID) {
	procMutex.Lock()
	defer procMutex.Unlock()
	if task, ok := processes[id]; ok {
		task.Process.Process.Kill()
	} else {
		log.Println("Attempting to kill task that does not exist", id)
	}
}

func Start() error {
	server := kissrpc.NewServer(COM_PORT)
	server.AddFunc("prepareTask", prepareTask)
	server.AddFunc("startTask", startTask)
	server.AddFunc("taskKill", taskKill)
	return server.Start()
}
