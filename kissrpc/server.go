package kissprc

import (
	"net/http"
	"reflect"
	"time"
)

type Server struct {
	methodTable map[string]reflect.Value
}

func (this *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {

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
}

func (this *Client) Call(name string, args []interface{}) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (this *Client) Call1(name string, args []interface{}) (interface{}, error) {
	return []interface{}{}, nil
}
func (this *Client) Call2(name string, args []interface{}) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (this *Client) CallError(name string, args []interface{}) (interface{}, error) {
	return nil, nil
}
