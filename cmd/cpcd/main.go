package main

import (
	"log"

	"github.com/bahusvel/ClusterPipe/common"
)

var nodes = map[string]*common.CPD{}

func main() {
	err := Start()
	if err != nil {
		log.Println(err)
	}
}
