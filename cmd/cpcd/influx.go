package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/bahusvel/ClusterPipe/common"
	"github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB     = "square_holes"
	username = "bubba"
	password = "bumblebeetuna"
)

var influxClient client.Client

func InfluxInit() {
	// Make client
	var err error
	influxClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}

}

func treatValue(valueTable *map[string]interface{}, value reflect.Value, root string) {
	if !value.CanAddr() {
		return
	}
	switch value.Kind() {
	case reflect.Struct:
		traverseStruct(valueTable, value.Interface(), root)
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			treatValue(valueTable, value, fmt.Sprintf("%s.%d", root, i))
		}
	default:
		(*valueTable)[root] = value.Interface()
	}
}

func traverseStruct(valueTable *map[string]interface{}, value interface{}, root string) {
	vtype := reflect.TypeOf(value)
	vval := reflect.ValueOf(value)
	for i := 0; i < vtype.NumField(); i++ {
		structField := vtype.Field(i)
		fieldValue := vval.Field(i)
		treatValue(valueTable, fieldValue, root+"."+structField.Name)
	}
}

func TraverseParamTree(value interface{}) map[string]interface{} {
	valMap := map[string]interface{}{}
	traverseStruct(&valMap, value, "")
	return valMap
}

func LogUpdate(cpd string, status common.CPDStatus) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Create a point and add to batch
	tags := map[string]string{"host": cpd}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, status.LastCheckin)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(pt)

	// Write the batch
	influxClient.Write(bp)
}
