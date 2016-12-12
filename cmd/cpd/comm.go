package main

import (
	"log"

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

func receiveTask(task common.Task) error {
	log.Println("Received connection from CPD")
	err := Run(task)
	if err != nil {
		return err
	}
	return nil
}

func taskStat(id common.TaskID) common.Task {
	procMutex.RLock()
	defer procMutex.RUnlock()
	return processes[id]
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
	server.AddFunc("receiveTask", receiveTask)
	server.AddFunc("taskStat", taskStat)
	server.AddFunc("taskKill", taskKill)
	return server.Start()
}
