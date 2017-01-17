package main

import (
	"io"
	"log"
	"os"
	"reflect"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/bahusvel/ClusterPipe/kissrpc"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var clusterNodes []common.CPD
var clusterAddress = ""
var controller *kissrpc.Client

func init() {
	kissrpc.RegisterType([]*common.CPD{})
	kissrpc.RegisterType(common.Task{})
}

func checkCluster() {
	if clusterAddress == "" {
		log.Fatalln("You must specify the address of the cluster")
	}
	var err error
	controller, err = kissrpc.NewClient(clusterAddress)
	if err != nil {
		log.Fatalln("Failed to connect to cluster controller", err)
	}
	_, err = controller.Call("ping")
	if err != nil {
		log.Fatalln("Failed to connect to cluster controller", err)
	}
}

func fieldNames(value interface{}) []string {
	v := reflect.TypeOf(value)
	fieldNames := []string{}
	for i := 0; i < v.NumField(); i++ {
		fieldNames = append(fieldNames, v.Field(i).Name)
	}
	return fieldNames
}

func fieldValues(value interface{}) []string {
	v := reflect.ValueOf(value)
	fieldValues := []string{}
	for i := 0; i < v.NumField(); i++ {
		fieldValues = append(fieldValues, v.Field(i).String())
	}
	return fieldValues
}

func NodeStatus(c *cli.Context) error {
	checkCluster()
	nodes, err := controller.Call1("getNodes")
	if err != nil {
		log.Println("Error getting nodes", err)
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fieldNames(common.CPD{}))
	for _, node := range nodes.([]*common.CPD) {
		table.Append(fieldValues(*node))
	}
	table.Render()
	return nil
}

func Exec(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.NewExitError("Command to execute was not specified", -1)
	}
	task := common.Task{Command: c.Args().First(), Args: []string(c.Args())[1:]}
	tskInter, err := controller.Call1("prepareTask", task)
	if err != nil {
		return err
	}
	task = tskInter.(common.Task)
	pipes, err := common.RemotePipe(task.Node, task.TID)
	if err != nil {
		return err
	}
	_, err = kissrpc.SingleCall(task.Node, "startTask", task.TID)
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, pipes.Stdout)
	go io.Copy(os.Stderr, pipes.Stderr)
	_, err = io.Copy(pipes.Stdin, os.Stdin)

	return err
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "ps",
		},
		{
			Name:   "nodes",
			Action: NodeStatus,
		},
		{
			Name: "exec",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "sched, s",
				},
				cli.IntSliceFlag{
					Name:  "stdio, i",
					Value: &cli.IntSlice{0, 1, 2},
				},
			},
			Action: Exec,
		},
		{
			Name: "kill",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "job, j",
				},
			},
		},
		{
			Name:  "checkpoint",
			Usage: "Perform manual service checkpoint to all nodes in cluster",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "master, m",
				},
			},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "cluster, c",
			Destination: &clusterAddress,
		},
	}

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}
	app.Run(os.Args)
}
