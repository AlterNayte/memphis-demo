package main

import (
	"fmt"
	"memphis-demo/testpb"
	"os"

	"github.com/memphisdev/memphis.go"
)

func main() {
	conn, err := memphis.Connect("localhost", "go.user", memphis.Password("PASSWORD"))
	if err != nil {
		fmt.Println("Connection failed: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	p, err := conn.CreateProducer("test-station", "go.producer")
	if err != nil {
		fmt.Println("Producer failed: %v", err)
		os.Exit(1)
	}

	hdrs := memphis.Headers{}
	hdrs.New()
	err = hdrs.Add("key", "value")

	if err != nil {
		fmt.Errorf("Header failed: %v", err)
		os.Exit(1)
	}

	msg := testpb.Test{
		Field1: "Hello",
		Field2: "World",
		Field3: 42,
	}

	err = p.Produce(&msg, memphis.MsgHeaders(hdrs))

	if err != nil {
		fmt.Errorf("Produce failed: %v", err)
		os.Exit(1)
	}
}
