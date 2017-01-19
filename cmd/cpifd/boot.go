package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pin/tftp"
)

const (
	CONTROLLER_IP   = "192.168.7.28"
	TFTPROOT        = "/tftpboot/original/"
	NFSROOT         = "/nfs/"
	CMDLINE         = "dwc_otg.lpm_enable=0 console=serial0,115200 console=tty1 root=/dev/nfs nfsroot=%s:%s,udp,vers=3,rsize=32768 rw ip=dhcp rootwait elevator=deadline smsc95xx.turbo_mode=Y"
	ADD_NODE_SCRIPT = `cd /nfs
	btrfs subvol snap reference $1
	cd $1
	rm etc/ssh/ssh_host_*
	ssh-keygen -t dsa -N "" -f etc/ssh/ssh_host_dsa_key
	ssh-keygen -t rsa -N "" -f etc/ssh/ssh_host_rsa_key
	ssh-keygen -t ecdsa -N "" -f etc/ssh/ssh_host_ecdsa_key
	echo "$1" > etc/hostname
	echo "127.0.1.1 $1" >> etc/hosts
	echo "/nfs/$1 *(rw,async,no_subtree_check,no_root_squash)" >> /etc/exports
	exportfs -ra`
)

func checkAndRegister(serialNumber string) {
	if _, err := os.Stat(NFSROOT + serialNumber); err == nil {
		log.Printf("Node %s is already registered\n", serialNumber)
		return
	}
	log.Println("Registering node", serialNumber)
	cmd := exec.Command("bash", "-c", ADD_NODE_SCRIPT, "add_node", serialNumber)
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to register node", serialNumber, err)
	}
}

func readHandler(filename string, rf io.ReaderFrom) error {
	remoteAddr := rf.(tftp.OutgoingTransfer).RemoteAddr().IP
	log.Printf("%s wants %s\n", remoteAddr, filename)
	parts := strings.Split(filename, "/")
	var reader io.Reader
	if len(parts) > 1 {
		serialNumber := parts[0]
		// Serial numbered path
		if parts[1] == "cmdline.txt" {
			checkAndRegister(serialNumber)
			cmdline_buffer := bytes.NewBufferString(fmt.Sprintf(CMDLINE, CONTROLLER_IP, NFSROOT+serialNumber))

			rf.(tftp.OutgoingTransfer).SetSize(int64(cmdline_buffer.Len()))
			reader = cmdline_buffer
		} else {
			var err error
			reader, err = os.Open(TFTPROOT + parts[1])
			if err != nil {
				log.Println(err)
				return err
			}
		}
	} else {
		var err error
		reader, err = os.Open(TFTPROOT + filename)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	n, err := rf.ReadFrom(reader)
	if err != nil {
		log.Printf("Error sending %s to %s %v\n", filename, remoteAddr, err)
		return err
	}
	if closer, ok := reader.(io.ReadCloser); ok {
		closer.Close()
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

func main() {
	StartTFTP()
}
