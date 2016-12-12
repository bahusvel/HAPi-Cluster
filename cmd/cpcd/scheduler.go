package main

import (
	"fmt"

	"github.com/bahusvel/ClusterPipe/common"
)

type PreparedJob struct {
}

var taskID = common.TaskID(0)

func ScheduleTask(task *common.Task) error {
	nodes := getNodes()
	if len(nodes) == 0 {
		return fmt.Errorf("Cluster does not have any nodes")
	}
	task.Node = nodes[int(taskID)%len(nodes)].Host
	return nil
}
