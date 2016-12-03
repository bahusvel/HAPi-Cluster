package main

import (
	"time"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/valyala/gorpc"
)

const (
	COM_PORT = "3334"
)

func ping() {

}

func getNodes() []*common.CPD {
	currentNodes := []*common.CPD{}
	for _, node := range nodes {
		currentNodes = append(currentNodes, node)
	}
	return currentNodes
}

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
	dispatch.AddFunc("getNodes", getNodes)
	dispatch.AddFunc("ping", ping)
	server := gorpc.NewTCPServer(":"+COM_PORT, dispatch.NewHandlerFunc())
	return server.Serve()
}
