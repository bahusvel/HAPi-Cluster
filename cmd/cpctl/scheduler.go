package main

const (
	CP_ROOT  = "/var/clusterpipe/"
	CP_PIPES = CP_ROOT + "pipes/"
)

type FIFO struct {
	Name       string
	References []string
}

func MkFIFO() error {
	return nil
}

func BridgeFIFO(in FIFO, toHost string) error {
	return nil
}

func Run(in *FIFO, command string, out *FIFO) error {
	return nil
}
