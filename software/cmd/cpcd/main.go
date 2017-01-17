package main

import (
	"log"
	"os"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/urfave/cli"
)

var nodes = map[string]*common.CPD{}
var useInflux = false

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "use-influx, i",
			Destination: &useInflux,
		},
	}
	app.Action = func(c *cli.Context) error {
		if useInflux {
			err := InfluxInit()
			if err != nil {
				return err
			}
			common.MonitorDiskUsage = true
			go common.StartStatMonitor(func(status common.CPDStatus) {
				err := InfluxInsert("controller", status)
				if err != nil {
					log.Println("Failed to insert controller status to influx", err)
				}
			})
		}
		err := Start()
		if err != nil {
			log.Println(err)
		}
		return nil
	}
	app.Run(os.Args)
}
