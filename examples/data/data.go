package main

import (
	"fmt"
	"log"

	"github.com/edwinwalela/africastalking-go/pkg/data"
)

func main() {
	client := &data.Client{
		ApiKey:    " ",
		Username:  "",
		IsSandbox: true,
	}

	dataRequest := &data.DataRequest{
		Username:    "",
		ProductName: "",
		Recipients: []data.Recipient{
			{PhoneNumber: "", Quantity: "", Unit: "", Validity: "", IsPromoBundle: "", MetaData: ""},
		},
	}

	response, err := client.SendMobileData(dataRequest)

	if err != nil {
		log.Fatalf("Error making request to Africa's Talking API %v", err)
	}

	fmt.Println(response)
}
