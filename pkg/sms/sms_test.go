package sms

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSendBulk(t *testing.T) {
	client := &Client{
		apiKey:    os.Getenv("AT_API_KEY"),
		username:  os.Getenv("AT_USERNAME"),
		isSandbox: true,
	}
	bulkRequest := &BulkRequest{
		To:            []string{"+254706496885"},
		Message:       "Hello AT",
		From:          "",
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}
	response, err := client.SendBulk(bulkRequest)
	if err != nil {
		t.Fatalf("bulk sms request failed: %s", err.Error())
	}
	fmt.Println(response)
}

func TestSendPremium(t *testing.T) {
	premiumRequest := &PremiumRequest{
		To:            []string{""},
		Message:       "Hello AT",
		From:          "",
		Keyword:       "",
		Enqueue:       true,
		LinkId:        "",
		RetryDuration: time.Hour,
		RequestId:     "",
	}
	_ = premiumRequest
}
