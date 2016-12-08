package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bahusvel/ClusterPipe/common"
)

func ReadJobFile(path string) (*common.Job, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	job := common.Job{}
	err = json.Unmarshal(data, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func TestJobEntry(t *testing.T) {
	job, err := ReadJobFile("../../docs/example_job.json")
	if err != nil {
		t.Error(err)
	}
	JobEntry(*job)
}
