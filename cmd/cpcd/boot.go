package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/pin/tftp"
)

func readHandler(filename string, rf io.ReaderFrom) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	return fmt.Errorf("Writes are administratively prohibited")
}

func StartTFTP() {
	// use nil in place of handler to disable read or write operations
	s := tftp.NewServer(readHandler, writeHandler)
	s.SetTimeout(5 * time.Second) // optional
	err := s.ListenAndServe(":69")
	if err != nil {
		log.Fatal(err)
	}
}
