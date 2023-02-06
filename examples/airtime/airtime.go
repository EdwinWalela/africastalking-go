package main

import (
	"fmt"
	"os"

	"github.com/edwinwalela/africastalking-go/pkg/airtime"
)

func main() {
	// Define Africa's Talking Airtime client
	client := airtime.Client{
		Username:  os.Getenv("AT_USERNAME"),
		ApiKey:    os.Getenv("AT_API_KEY"),
		IsSandbox: true,
	}

	// Define a recipient to be topped up with airtime
	recipient := airtime.Recipient{
		PhoneNumber: "+25470000000001",
		Amount:      10.00,
		Currency:    airtime.KES,
	}

	// Define the request body for the top up request
	request := &airtime.Request{
		Recipients: []airtime.Recipient{recipient},
	}

	// Send the top up request
	response, err := client.Send(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
