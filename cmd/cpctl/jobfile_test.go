package main

import (
	"log"
	"testing"
)

func TestReadJobFile(t *testing.T) {
	job, err := ReadJobFile("../../docs/example_job.json")
	if err != nil {
		t.Error(err)
	}
	log.Println(job)
}
