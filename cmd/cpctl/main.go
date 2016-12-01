package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "ps",
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
