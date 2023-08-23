package data

import "testing"

func TestData(t *testing.T) {
	client := &Client{
		ApiKey:    "",
		Username:  "",
		IsSandbox: true,
	}

	dataRequest := &DataRequest{
		Username:    "",
		ProductName: "",
		Recipients: []Recipient{
			{PhoneNumber: "", Quantity: "", Unit: "", Validity: "", IsPromoBundle: "", MetaData: ""},
		},
	}

	response, err := client.SendMobileData(dataRequest)

	if err != nil {
		t.Fatalf("Send data request failed: %s", err.Error())
	}
	status := response.Entries[0].Status
	if status != "Success" {
		t.Fatalf("expected status = 'success' got status = '%s'", status)
	}
}
