package main

import (
	"time"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/valyala/gorpc"
)

const (
	COM_PORT = "3334"
)

func registerCPD(node *common.CPD) {
	nodes[node.Host] = node
}

func updateCPD(cpd string, status *common.CPDStatus) {
	status.LastCheckin = time.Now()
	nodes[cpd].CurrentStatus = status
}

func Start() error {
	dispatch := gorpc.NewDispatcher()
	dispatch.AddFunc("registerCPD", registerCPD)
	dispatch.AddFunc("updateCPD", updateCPD)
	server := gorpc.NewTCPServer(":"+COM_PORT, dispatch.NewHandlerFunc())
	return server.Serve()
}
