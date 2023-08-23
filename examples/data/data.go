package main

import (
	"fmt"
	"log"

	"github.com/edwinwalela/africastalking-go/pkg/data"
)

func main() {
	var apiKey = "4be4c57c242eefbcafe95f690f77d4ec1f2ec027df8c36320937006f8ddc3af2"
	fmt.Println(apiKey)
	client := &data.Client{
		ApiKey:    apiKey,
		Username:  "DanielSogbey",
		IsSandbox: true,
	}

	dataRequest := &data.DataRequest{
		Username:    client.Username,
		ProductName: "Open Source Software",
		Recipients: []data.Recipient{
			{PhoneNumber: "+233558159629"},
		},
		Quantity:      "1",
		Unit:          "MB",
		Validity:      "Day",
		IsPromoBundle: "false",
	}

	response, err := client.SendMobileData(dataRequest)

	if err != nil {
		log.Fatalf("Error making request to Africa's Talking API %v", err)
	}

	fmt.Println(response)
}
