package data

import (
	"testing"
)

func TestSendData(t *testing.T) {
	client := &Client{
		ApiKey:    "f0d2459e095f1bc43aaedcfb3dff0a45ff9c32e16794a9c21b170fbb6df1a657",
		Username:  "DanielSogbey",
		IsSandbox: true,
	}

	dataRequest := &Request{
		Username:    "DanielSogbey",
		ProductName: "Open Source Software",
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
