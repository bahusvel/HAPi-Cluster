package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

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
