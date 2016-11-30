package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"syscall"
)

const (
	CONN_PORT      = "3333"
	PIPE_NAME_SIZE = 64
	PIPE_MODE      = 0666
)

type PipeRequest struct {
	PipeName [PIPE_NAME_SIZE]byte
}

func RequestPipe(host string, pipename string) error {
	if len(pipename) >= PIPE_NAME_SIZE {
		return fmt.Errorf("Pipename too long")
	}
	pipeName := CP_PIPES + string(pipename)
	err := syscall.Mkfifo(pipeName, PIPE_MODE)
	if err != nil {
		log.Println("Error creating pipe", err)
		return err
	}
	pipe, err := os.Open(pipeName)
	if err != nil {
		log.Println("Error opening pipe", err)
		return err
	}
	request := PipeRequest{}
	copy(request.PipeName[:], []byte(pipename))
	conn, err := net.Dial("tcp", host+":"+CONN_PORT)
	if err != nil {
		log.Println("Error creating tcp connection", err)
		os.Remove(pipeName)
		return err
	}
	err = binary.Write(conn, binary.BigEndian, request)
	if err != nil {
		log.Println("Error writing pipe request", err)
		os.Remove(pipeName)
		return err
	}
	go io.Copy(conn, pipe)
	_, err = io.Copy(pipe, conn)

	return err
}

/*
Netpipe handshake:
1) C:Send request for pipe
2) S:Create pipe
3) S:Send ACK to C
4) S:Bridge pipe to tcp socket
5) C,S: Communicate as normal
*/
func handleRequest(conn net.Conn) {
	defer conn.Close()
	// Make a buffer to hold incoming data.
	request := PipeRequest{}
	// Read the incoming connection into the buffer.
	err := binary.Read(conn, binary.BigEndian, &request)
	if err != nil {
		log.Println("Error reading request", err)
		conn.Write([]byte{1})
		return
	}
	n := bytes.IndexByte(request.PipeName[:], 0)
	pipeName := CP_PIPES + string(request.PipeName[:n])
	err = syscall.Mkfifo(pipeName, PIPE_MODE)
	if err != nil {
		log.Println("Error creating pipe", err)
		conn.Write([]byte{1})
		return
	}
	defer os.Remove(pipeName)
	pipe, err := os.Open(pipeName)
	if err != nil {
		log.Println("Error opening pipe", err)
		conn.Write([]byte{1})
		return
	}
	conn.Write([]byte{0})

	go io.Copy(conn, pipe)
	io.Copy(pipe, conn)
}

func RunServer() {
	l, err := net.Listen("tcp", ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on :" + CONN_PORT)
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
