package main

import (
	"github.com/bahusvel/ClusterPipe/common"
	"github.com/valyala/gorpc"
)

const (
	COM_PORT = "3334"
)

func registerCPD(node string) {
	nodes = append(nodes, common.CPD{Host: node})
}

func Start(internalip string) error {
	dispatch := gorpc.NewDispatcher()
	dispatch.AddFunc("registerCPD", registerCPD)
	server := gorpc.NewTCPServer(":"+COM_PORT, dispatch.NewHandlerFunc())
	return server.Serve()
}
