package main

import (
	"log"
	"net/http"

	"github.com/bahusvel/ClusterPipe/common/httputil"
)

const (
	COM_PORT = "3334"
)

func receiveTask(w http.ResponseWriter, r *http.Request) {
	log.Println("Received connection from CPD")
	httputil.WriteSimpleResponse(w, "OK")
}

func makeFifo(w http.ResponseWriter, r *http.Request) {
	err := MkFIFO()
	if err != nil {
		httputil.WriteSimpleResponse(w, "Error - "+err.Error())
		return
	}
	httputil.WriteSimpleResponse(w, "OK")
}

func Start(internalip string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks/receive", receiveTask)
	mux.HandleFunc("/fifo/make", makeFifo)
	server := &http.Server{Addr: ":" + COM_PORT, Handler: mux}
	log.Println("Server running on ", ":"+COM_PORT)
	return server.ListenAndServe()
}
