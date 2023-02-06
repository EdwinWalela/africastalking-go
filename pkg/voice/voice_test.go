package voice

import (
	"fmt"
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
		From: "+254706496885",
		To:   []string{"+254706496885"},
	}

	response, err := client.Call(request)
	if err != nil {
		fmt.Println(err)
	}
	_ = response
}
