package main

import (
	"fmt"

	"github.com/bahusvel/ClusterPipe/common"
)

var schedulers = map[string]Scheduler{
	"rr":       &RoundRobin{},
	"lowest":   NoneScheduler{},
	"specific": NoneScheduler{},
	"same":     NoneScheduler{},
}

type Scheduler interface {
	Schedule(task *common.Task, schedulingArgs []interface{}) error
}

var taskIDIncrement = common.TaskID(0)

type RoundRobin struct {
	counter int
}

func (this *RoundRobin) Schedule(task *common.Task, schedulingArgs []interface{}) error {
	nodes := getNodes()
	if len(nodes) == 0 {
		return fmt.Errorf("Cluster does not have any nodes")
	}
	task.Node = nodes[this.counter%len(nodes)].Host
	this.counter++
	task.TID = taskIDIncrement
	taskIDIncrement++
	return nil
}

type NoneScheduler struct {
}

func (this NoneScheduler) Schedule(task *common.Task, schedulingArgs []interface{}) error {
	return fmt.Errorf("This method is not yet implemented")
}
