package common

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
	PIPE_MODE      = 0777
	CONN_PORT      = "3333"
	PIPE_NAME_SIZE = 64
)

type PipeRequest struct {
	PipeName [PIPE_NAME_SIZE]byte
}

type FIFO struct {
	Job        string
	Node       string
	GivenName  string
	References []string
}

func (this FIFO) MakeLocal() error {
	err := syscall.Mkfifo(this.AsArgument(), PIPE_MODE)
	if err != nil {
		return err
	}
	return nil
}

func (this FIFO) MakeRemote() error {
	if len(this.AsArgument()) >= PIPE_NAME_SIZE {
		return fmt.Errorf("Pipename too long")
	}
	err := this.MakeLocal()
	if err != nil {
		log.Println("Error creating pipe", err)
		return err
	}
	pipeFile, err := this.Open()
	if err != nil {
		log.Println("Error opening pipe", err)
		return err
	}
	request := PipeRequest{}
	copy(request.PipeName[:], []byte(this.AsArgument()))
	conn, err := net.Dial("tcp", this.Node+":"+CONN_PORT)
	if err != nil {
		log.Println("Error creating tcp connection", err)
		os.Remove(this.AsArgument())
		return err
	}
	err = binary.Write(conn, binary.BigEndian, request)
	if err != nil {
		log.Println("Error writing pipe request", err)
		os.Remove(this.AsArgument())
		return err
	}
	go io.Copy(conn, pipeFile)
	_, err = io.Copy(pipeFile, conn)

	return err
}

func (this FIFO) AsArgument() string {
	return fmt.Sprintf("%s%s/%s", CP_PIPES, this.Job, this.GivenName)
}

func (this FIFO) Open() (*os.File, error) {
	return os.Open(this.AsArgument())
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
	pipeName := string(request.PipeName[:n])
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

func RunPipeServer() {
	l, err := net.Listen("tcp", ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Fifo server listenting on :" + CONN_PORT)
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
