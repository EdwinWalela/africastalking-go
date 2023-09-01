package data

import (
	"os"
	"testing"
)

func TestSendData(t *testing.T) {
	client := &Client{
		ApiKey:    os.Getenv("AT_API_KEY"),
		Username:  os.Getenv("AT_USERNAME"),
		IsSandbox: true,
	}

	dataRequest := &Request{
		Username:    os.Getenv("AT_USERNAME"),
		ProductName: os.Getenv("AT_PRODUCT_NAME"),
		Recipients: []Recipient{
			{PhoneNumber: "+233558159629", Quantity: 2, Unit: "MB", Validity: "Day", IsPromoBundle: "true", MetaData: ""},
		},
	}

	response, err := client.Send(dataRequest)

	if err != nil {
		t.Fatalf("Send data request failed: %s", err.Error())
	}
	status := response.Entries[0].Status
	if status != "Success" {
		t.Fatalf("expected status = 'success' got status = '%s'", status)
	}
}
