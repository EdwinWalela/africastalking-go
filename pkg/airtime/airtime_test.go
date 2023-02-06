package airtime

import (
	"fmt"
	"os"
	"testing"
)

func TestSendAirtime(t *testing.T) {
	client := Client{
		Username:  os.Getenv("AT_USERNAME"),
		ApiKey:    os.Getenv("AT_API_KEY"),
		IsSandbox: true,
	}

	recipient := Recipient{
		PhoneNumber: "+254706496885",
		Amount:      10,
		Currency:    KES,
	}

	request := &Request{
		Recipients: []Recipient{recipient},
	}

	response, err := client.Send(request)
	if err != nil {
		t.Fatalf("airtime request failed: %s", err.Error())
	}
	fmt.Println(response)
}
