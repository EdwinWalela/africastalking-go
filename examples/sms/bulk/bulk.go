package main

import (
	"fmt"

	"github.com/edwinwalela/africastalking-go/pkg/sms"
)

func main() {
	client := &sms.Client{}
	request := &sms.BulkRequest{}

	response, err := client.SendBulk(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
