package main

import (
	"os/exec"

	"github.com/bahusvel/ClusterPipe/common"
)

var localNode string

func ScheduleTasks(tasks []common.PreparedTask) error {
	for _, task := range tasks {
		args := []string{}
		for _, arg := range task.Args {
			if fifo, ok := arg.(common.FIFO); ok {
				if err := fifo.Check(); err != nil {
					return err
				}
				args = append(args, arg.AsArgument())
			}
		}

		cmd := exec.Command(task.Command, args...)
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

		err = Run(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func Run(command *exec.Cmd) error {
	return nil
}

func MkFIFO() error {
	return nil
}

func main() {

}
