package kissrpc

import (
	"log"
	"testing"
)

type Message struct {
	Text   string
	Number int
}

func TestSimpleServer(t *testing.T) {
	server := NewServer()
	server.AddFunc("Test", func(message Message) {
		log.Println("Hello!", message.Text)
	})
	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func TestSimpleCall(t *testing.T) {
	client, err := NewClient("127.0.0.1:8080")
	if err != nil {
		t.Error(err.Error())
	}
	client.RegisterType(Message{})
	_, err = client.Call("Test", Message{Text: "Hi", Number: 1})
	if err != nil {
		t.Error(err.Error())
	}
}
