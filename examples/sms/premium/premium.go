package main

import (
	"fmt"

	"github.com/edwinwalela/africastalking-go/pkg/sms"
)

func main() {
	client := &sms.Client{}
	request := &sms.PremiumRequest{}

	response, err := client.SendPremium(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Message)
}
