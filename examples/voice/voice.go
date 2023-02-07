package main

import (
	"fmt"
	"os"

	"github.com/edwinwalela/africastalking-go/pkg/voice"
)

func main() {
	// Define Africa's Talking Voice client
	client := voice.Client{
		Username:  os.Getenv("AT_USERNAME"),
		ApiKey:    os.Getenv("AT_API_KEY"),
		IsSandbox: true,
	}

	// Define the request body for the call request
	request := &voice.Request{
		From: os.Getenv("AT_CALLER_ID"),
		To:   []string{"+254700000001"},
	}

	// Initiate the call
	response, err := client.Call(request)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

}
