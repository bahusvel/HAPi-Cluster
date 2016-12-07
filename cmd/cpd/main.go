package main

import (
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/bahusvel/ClusterPipe/kissrpc"
	"github.com/urfave/cli"
)

var controllerAddress string
var controller *kissrpc.Client
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
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "controller, c",
			Destination: &controllerAddress,
		},
		cli.StringFlag{
			Name:        "ip, i",
			Destination: &thisCPD.Host,
		},
	}

	app.Action = func(c *cli.Context) error {
		if controllerAddress == "" {
			return cli.NewExitError("You must specify controller address -c", -1)
		}
		if thisCPD.Host == "" {
			return cli.NewExitError("You must specify local address -i", -1)
		}

		go common.RunPipeServer()
		go Start()
		err := GatherHostInfo()
		if err != nil {
			return err
		}
		controller, err = kissrpc.NewClient(controllerAddress)
		if err != nil {
			return err
		}
		_, err = controller.Call("registerCPD", &thisCPD)
		if err != nil {
			return err
		}
		for {
			time.Sleep(1 * time.Second)
		}
	}

	app.Run(os.Args)

}
