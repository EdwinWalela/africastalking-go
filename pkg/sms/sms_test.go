package sms

import (
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
	status := response.Recipients[0].Status
	if status != "Success" {
		t.Fatalf("expected status = 'success' got status = '%s'", status)
	}
}

func TestSendPremium(t *testing.T) {
	client := &Client{
		apiKey:    os.Getenv("AT_API_KEY"),
		username:  os.Getenv("AT_USERNAME"),
		isSandbox: true,
	}
	premiumRequest := &PremiumRequest{
		To:            []string{"+254706496885"},
		Message:       "Hello AT",
		From:          "",
		Keyword:       "",
		Enqueue:       true,
		LinkId:        "",
		RetryDuration: time.Hour,
		RequestId:     "",
	}
	response, err := client.SendPremium(premiumRequest)
	if err != nil {
		t.Fatalf("bulk sms request failed: %s", err.Error())
	}
	status := response.Recipients[0].Status
	if status != "Success" {
		t.Fatalf("expected status = 'success' got status = '%s'", status)
	}
}
