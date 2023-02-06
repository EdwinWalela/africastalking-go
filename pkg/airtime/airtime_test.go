package airtime

import (
	"os"
	"testing"
)

func TestSendAirtime(t *testing.T) {
	client := Client{
		Username: os.Getenv("AT_USERNAME"),
		ApiKey:   os.Getenv("AT_APIKEY"),
	}

	recipient := Recipient{
		PhoneNumber: "+254706496885",
		Amount:      10,
		Currency:    "KES",
	}

	request := &Request{
		Recipients: []Recipient{recipient, recipient},
	}

	client.Send(request)
}
