package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bahusvel/ClusterPipe/common"
)

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	request := common.PipeRequest{}
	// Read the incoming connection into the buffer.
	err := binary.Read(conn, binary.BigEndian, &request)
	if err != nil {
		log.Println("Error reading request", err)
		conn.Write([]byte{1})
		return
	}
	task, ok := processes[request.TID]
	if !ok {
		log.Println("request stdio to non existant task", request.TID)
		conn.Close()
	}
	if request.Flags&common.STDIN != 0 {
		task.Stdio.Stdin = conn
	}
	if request.Flags&common.STDOUT != 0 {
		task.Stdio.Stdout = conn
	}
	if request.Flags&common.STDERR != 0 {
		task.Stdio.Stderr = conn
	}

}

func RunPipeServer() {
	l, err := net.Listen("tcp", ":"+common.CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Fifo server listenting on :" + common.CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
