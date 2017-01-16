package common

import (
	"encoding/binary"
	"log"
	"net"
)

const (
	PIPE_MODE      = 0777
	CONN_PORT      = "3333"
	PIPE_NAME_SIZE = 64
)

const (
	STDIN  = 1 << iota
	STDOUT = 1 << iota
	STDERR = 1 << iota
)

type PipeRequest struct {
	TID   TaskID
	Flags uint8
}

func RemotePipe(node string, task TaskID) (*TaskPipes, error) {
	pipes := &TaskPipes{}
	request := PipeRequest{TID: task}
	if request.Flags == 0 {
		return pipes, nil
	}
	stdinout, err := net.Dial("tcp", node+":"+CONN_PORT)
	if err != nil {
		log.Println("Error creating tcp connection", err)
		return pipes, err
	}
	request.Flags = STDIN | STDOUT
	err = binary.Write(stdinout, binary.BigEndian, request)
	if err != nil {
		log.Println("Error writing pipe request", err)
		return pipes, err
	}
	pipes.Stdin = stdinout
	pipes.Stdout = stdinout
	stderr, err := net.Dial("tcp", node+":"+CONN_PORT)
	if err != nil {
		log.Println("Error creating tcp connection", err)
		return pipes, err
	}
	request.Flags = STDERR
	err = binary.Write(stderr, binary.BigEndian, request)
	if err != nil {
		log.Println("Error writing pipe request", err)
		return pipes, err
	}
	pipes.Stderr = stderr
	return pipes, err
}
