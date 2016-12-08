package main

import (
	"log"
	"testing"

	"github.com/bahusvel/ClusterPipe/common"
)

func TestWeirdMapThingy(t *testing.T) {
	cpdStatus := common.CPDStatus{}
	mapThingy := TraverseParamTree(cpdStatus)
	log.Println(mapThingy)
}
