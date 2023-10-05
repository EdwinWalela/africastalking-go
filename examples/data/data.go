package main

import (
	"fmt"
	"log"
	"os"

	"github.com/edwinwalela/africastalking-go/pkg/data"
)

func main() {
	client := &data.Client{
		ApiKey:    os.Getenv("AT_API_KEY"),
		Username:  os.Getenv("AT_USERNAME"),
		IsSandbox: true,
	}

	dataRequest := &data.Request{
		Username:    os.Getenv("AT_USERNAME"),
		ProductName: os.Getenv("AT_PRODUCT_NAME"),
		Recipients: []data.Recipient{
			{PhoneNumber: "+233558159629", Quantity: 2, Unit: "MB", Validity: "Day", IsPromoBundle: "true"},
		},
	}

	response, err := client.Send(dataRequest)

	if err != nil {
		log.Fatalf("Error making request to Africa's Talking API %v", err)
	}

	fmt.Println(response)
}
