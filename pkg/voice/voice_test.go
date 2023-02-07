package voice

import (
	"os"
	"testing"
)

func TestCall(t *testing.T) {
	client := Client{
		ApiKey:    os.Getenv("AT_API_KEY"),
		Username:  os.Getenv("AT_USERNAME"),
		IsSandbox: true,
	}

	request := &Request{
		From: os.Getenv("AT_CALLER_ID"),
		To:   []string{"+254700000001"},
	}

	response, err := client.Call(request)
	if err != nil {
		t.Fatalf("failed to initiate call: %s", err.Error())
	}

	if response.ErrorMessage != "None" {
		t.Fatalf("expected errorMessage='none' got errorMessage='%s'", response.ErrorMessage)
	}
}
