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
	request := &voice.CallTransferRequest{
		SessionId:    "session-id",
		PhoneNumber:  "+254700000001",
		CallLeg:      "callee",
		HoldMusicUrl: "https://my-server.com/audio/hold-music.mp3",
	}

	// Initiate the call transfer
	response, err := client.Transfer(request)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)

}
