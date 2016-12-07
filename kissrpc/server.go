package kissrpc

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"
)

var registeredTypes = map[string]struct{}{
	"string": {},
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
	log.Println("Calling", callRequest.Name)
	if method, ok = this.methodTable[callRequest.Name]; !ok {
		log.Println("Requested method not found", callRequest.Name)
		return
	}

	arguments := []reflect.Value{}
	for _, arg := range callRequest.Args {
		log.Println("Got arg", arg)
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
	log.Printf("Function %s has type %s\n", name, funcType.String())
	for i := 0; i < funcType.NumIn(); i++ {
		registerType(funcType.In(i))
	}
	for i := 0; i < funcType.NumOut(); i++ {
		registerType(funcType.Out(i))
	}
	this.methodTable[name] = val
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

type Client struct {
	serverUrl string
}

func NewClient(address string) (*Client, error) {
	client := Client{serverUrl: "http://" + address}
	return &client, nil
}

func registerType(inType reflect.Type) {
	if _, ok := registeredTypes[inType.String()]; !ok {
		log.Println("Registering type", inType.String())
		gob.Register(reflect.New(inType).Interface())
		registeredTypes[inType.String()] = struct{}{}
	}
}

func RegisterType(regType interface{}) {
	gob.Register(regType)
}

func (this Client) Call(name string, args ...interface{}) ([]interface{}, error) {
	var requestBuffer bytes.Buffer
	enc := gob.NewEncoder(&requestBuffer)
	err := enc.Encode(call{Name: name, Args: args})
	if err != nil {
		return []interface{}{}, err
	}
	response, err := http.Post(this.serverUrl, "application/data", &requestBuffer)
	if err != nil {
		return []interface{}{}, err
	}
	retValues := callReturn{}
	dec := gob.NewDecoder(response.Body)
	err = dec.Decode(&retValues)
	if err != nil {
		return []interface{}{}, err
	}
	return retValues.ReturnValues, nil
}

func (this Client) Call1(name string, args ...interface{}) (interface{}, error) {
	var rets []interface{}
	var err error
	rets, err = this.Call(name, args...)
	if err != nil {
		return []interface{}{}, err
	}
	if len(rets) != 1 {
		return []interface{}{}, fmt.Errorf("Unexpected return values for %s expected %d got %d", name, 1, len(rets))
	}
	return rets[0], err
}
func (this Client) Call2(name string, args ...interface{}) (interface{}, interface{}, error) {
	var rets []interface{}
	var err error
	rets, err = this.Call(name, args)
	if len(rets) != 2 {
		return []interface{}{}, []interface{}{}, fmt.Errorf("Unexpected return values for %s expected %d got %d", name, 2, len(rets))
	}
	return rets[0], rets[1], err
}

func (this Client) CallError(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}
