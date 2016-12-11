package main

import (
	"fmt"

	"github.com/bahusvel/ClusterPipe/common"
)

type PreparedJob struct {
}

var counter = 0

func ScheduleTask(job common.Job, task *common.PreparedTask) error {
	nodes := getNodes()
	if len(nodes) == 0 {
		return fmt.Errorf("Cluster does not have any nodes")
	}
	task.Node = nodes[counter%len(nodes)].Host
	return nil
}

func JobEntry(job common.Job) error {
	for _, task := range job.Flow {
		preparedTask := common.PreparedTask{Job: job.Name, Command: task.Command}
		err := ScheduleTask(job, &preparedTask)
		if err != nil {
			return err
		}
		arguments := []interface{}{}
		for _, arg := range task.Args {
			if arg[0] == '$' {
				fifoArg := common.FIFO{Job: job.Name, GivenName: arg[1:], Node: preparedTask.Node}
				arguments = append(arguments, fifoArg)
			} else {
				arguments = append(arguments, arg)
			}
		}
	}
	return nil
}
