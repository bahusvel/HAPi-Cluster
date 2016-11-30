package common

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

type StringArgument string
type Argument interface {
	AsArgument() string
}

type PreparedTask struct {
	Command string
	Args    []Argument
	Stdin   *FIFO
	Stdout  *FIFO
	Stderr  *FIFO
}

type Job struct {
	Owner string
	Name  string
	Quota ResourceQuota
	Flow  []Task
}
