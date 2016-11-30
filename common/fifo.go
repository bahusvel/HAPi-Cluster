package common

import (
	"fmt"
	"os"
)

type FIFO struct {
	Job        string
	GivenName  string
	Node       string
	References []string
}

func (this FIFO) AsArgument() string {
	return fmt.Sprintf("%s%s/%s", CP_PIPES, this.Job, this.GivenName)
}

func (this FIFO) Check() error {
	return nil
}

func (this FIFO) Open() (*os.File, error) {
	return nil, nil
}
