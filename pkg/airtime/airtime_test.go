package airtime

import (
	"os"
	"testing"
)

func TestSendAirtime(t *testing.T) {
	client := Client{
		Username:  os.Getenv("AT_USERNAME"),
		ApiKey:    os.Getenv("AT_API_KEY"),
		IsSandbox: true,
	}
	airtimeAmount := 10.00
	recipient1Phone := "+254700000001"
	recipient2Phone := "+254700000002"

	recipient1 := Recipient{
		PhoneNumber: recipient1Phone,
		Amount:      airtimeAmount,
		Currency:    KES,
	}

	recipient2 := Recipient{
		PhoneNumber: recipient2Phone,
		Amount:      airtimeAmount,
		Currency:    KES,
	}

	request := &Request{
		Recipients: []Recipient{recipient2, recipient1},
	}
	response, err := client.Send(request)
	expectedTotalAmount := float64(len(request.Recipients)) * airtimeAmount

	if err != nil {
		t.Fatalf("airtime request failed: %s", err.Error())
	}
	if response.ErrorMessage != "None" {
		t.Fatalf("expected errorMessage='None' got errorMessage='%s'", response.ErrorMessage)
	}
	if response.TotalAmount != expectedTotalAmount {
		t.Fatalf("expected totalAmount=%.2f got totalAmount=%.2f", expectedTotalAmount, response.TotalAmount)
	}
	if response.Responses[0].PhoneNumber != recipient1Phone {
		t.Fatalf("expected recipientPhone=%s got recipientPhone=%s", recipient1Phone, response.Responses[0].PhoneNumber)
	}
}
