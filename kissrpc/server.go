package kissrpc

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

const DEBUG = false

var registeredTypes = map[string]struct{}{
	"string": {},
	"float":  {},
	"int":    {},
	"int32":  {},
	"int64":  {},
	"error":  {},
}

func init() {
	RegisterType([]interface{}{})
}

type call struct {
	Name string
	Args []interface{}
}

type callReturn struct {
	ReturnValues []interface{}
	Error        error
}

type Server struct {
	address     string
	methodTable map[string]reflect.Value
}

func NewServer(address string) *Server {
	server := &Server{address: address}
	server.methodTable = make(map[string]reflect.Value)
	return server
}

func (this *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	callRequest := call{}
	dec := gob.NewDecoder(request.Body)
	err := dec.Decode(&callRequest)
	if err != nil {
		log.Println("Failed to read call request", err)
		return
	}
	var method reflect.Value
	var ok bool
	if DEBUG {
		log.Println("Calling", callRequest.Name)
	}
	if method, ok = this.methodTable[callRequest.Name]; !ok {
		log.Println("Requested method not found", callRequest.Name)
		return
	}

	arguments := []reflect.Value{}
	for _, arg := range callRequest.Args {
		arguments = append(arguments, reflect.ValueOf(arg))
	}

	retVals := []interface{}{}
	for _, retVal := range method.Call(arguments) {
		retVals = append(retVals, retVal.Interface())
	}

	enc := gob.NewEncoder(response)
	err = enc.Encode(callReturn{ReturnValues: retVals})
	if err != nil {
		log.Println("Failed to write call response", err)
		return
	}
}

func (this *Server) AddFunc(name string, function interface{}) {
	val := reflect.ValueOf(function)
	if val.Kind() != reflect.Func {
		panic(fmt.Errorf("%s is not a function", name))
	}
	funcType := val.Type()
	if DEBUG {
		log.Printf("Function %s has type %s\n", name, funcType.String())
	}
	for i := 0; i < funcType.NumIn(); i++ {
		registerType(funcType.In(i))
	}
	for i := 0; i < funcType.NumOut(); i++ {
		registerType(funcType.Out(i))
	}
	this.methodTable[name] = val
}

func (this *Server) AddService(service interface{}) {
	val := reflect.ValueOf(service)
	if val.Kind() != reflect.Struct {
		panic(fmt.Errorf("Supplied service is not a struct"))
	}
	serviceType := val.Type()
	for i := 0; i < serviceType.NumField(); i++ {
		field := serviceType.Field(i)
		if field.Type.Kind() != reflect.Func {
			panic(fmt.Errorf("Field %s of %s is not a function, only functions are supported", field.Name, serviceType.Name()))
		}
		this.AddFunc(serviceType.Name()+"."+field.Name, val.Field(i).Interface())
	}
}

func (this *Server) Start() error {
	s := &http.Server{
		Addr:           this.address,
		Handler:        this,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}
