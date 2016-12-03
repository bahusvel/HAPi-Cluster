package kissrpc

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"
	"reflect"
	"time"
)

type call struct {
	Name string
	Args []interface{}
}

type callReturn struct {
	ReturnValues []interface{}
	Error        error
}

type Server struct {
	methodTable map[string]reflect.Value
}

func NewServer() *Server {
	server := &Server{}
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
	if method, ok = this.methodTable[callRequest.Name]; !ok {
		log.Println("Requested method not found", callRequest.Name)
		return
	}
	arguments := []reflect.Value{}
	for _, arg := range callRequest.Args {
		arguments = append(arguments, reflect.ValueOf(arg))
	}
	_ = method.Call(arguments)
	enc := gob.NewEncoder(response)
	err = enc.Encode(callReturn{})
	if err != nil {
		log.Println("Failed to write call response", err)
		return
	}
}

func (this *Server) AddFunc(name string, function interface{}) {
	this.methodTable[name] = reflect.ValueOf(function)
}

func (this *Server) Start() error {
	s := &http.Server{
		Addr:           ":8080",
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

func (this Client) RegisterType(regType interface{}) {
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
	return []interface{}{}, nil
}

func (this Client) Call1(name string, args ...interface{}) (interface{}, error) {
	return []interface{}{}, nil
}
func (this Client) Call2(name string, args ...interface{}) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (this Client) CallError(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}
