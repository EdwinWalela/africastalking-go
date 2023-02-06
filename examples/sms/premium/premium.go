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

	// Define a request for the Premium SMS request
	premiumRequest := &sms.PremiumRequest{
		To:            []string{"+254706496885"},
		Message:       "Hello AT",
		From:          "",
		Keyword:       "",
		Enqueue:       true,
		LinkId:        "",
		RetryDuration: time.Hour,
		RequestId:     "",
	}

	// Send SMS to the defined recipients
	response, err := client.SendPremium(premiumRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
