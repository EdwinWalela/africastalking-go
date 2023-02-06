package main

import (
	"fmt"
	"os"
	"time"

	"github.com/edwinwalela/africastalking-go/pkg/sms"
)

func main() {
	// Define Africa's Talking SMS client
	client := &sms.Client{
		ApiKey:    os.Getenv("AT_API_KEY"),
		Username:  os.Getenv("AT_USERNAME"),
		IsSandbox: true,
	}

	// Define a request for the Bulk SMS request
	bulkRequest := &sms.BulkRequest{
		To:            []string{"+254706496885"},
		Message:       "Hello AT",
		From:          "",
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}

	// Send SMS to the defined recipients
	response, err := client.SendBulk(bulkRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
