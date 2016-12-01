package main

import (
	"os"
	"reflect"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var clusterNodes []common.CPD

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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fieldNames(common.CPD{}))
	for _, node := range clusterNodes {
		table.Append(fieldValues(node))
	}
	table.Render()
	return nil
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
			Name: "start",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "job, j",
				},
			},
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

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}
	app.Run(os.Args)
}
