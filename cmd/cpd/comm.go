package main

import (
	"log"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/bahusvel/ClusterPipe/kissrpc"
)

const (
	COM_PORT = ":3334"
)

func receiveTask(tasks []common.PreparedTask) error {
	log.Println("Received connection from CPD")
	for _, task := range tasks {
		err := Run(task)
		if err != nil {
			return err
		}
	}
	return nil
}

func jobStat(job string) []common.PreparedTask {
	procMutex.RLock()
	defer procMutex.RUnlock()
	return processes[job]
}

func jobKill(job string) {
	procMutex.Lock()
	defer procMutex.Unlock()
	for _, task := range processes[job] {
		task.Process.Process.Kill()
	}
}

func makeFifo(pipe common.FIFO) error {
	var err error
	if pipe.Node == thisCPD.Host {
		err = pipe.MakeLocal()
	} else {
		err = pipe.MakeRemote()
	}
	if err != nil {
		log.Println("Error creating pipe", err)
		return err
	}
	return nil
}

func Start() error {
	server := kissrpc.NewServer(COM_PORT)
	server.AddFunc("receiveTask", receiveTask)
	server.AddFunc("jobStat", jobStat)
	server.AddFunc("jobKill", jobKill)
	return server.Start()
}
