package kissrpc

import (
	"log"
	"testing"
	"time"
)

type Message struct {
	Text   string
	Number int
}

func TestSimpleServer(t *testing.T) {
	server := NewServer(":3334")
	server.AddFunc("Test", func(text string, number int) {
		log.Println("Hello!", text, number)
	})
	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(100 * time.Millisecond)
}

func TestSimpleCall(t *testing.T) {
	client, err := NewClient("127.0.0.1:3334")
	if err != nil {
		t.Error(err.Error())
	}
	RegisterType(Message{})
	_, err = client.Call("Test", "Test", 1)
	if err != nil {
		t.Error(err.Error())
	}
}
