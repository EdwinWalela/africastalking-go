package sms

import (
	"testing"
	"time"
)

func TestSendBulk(t *testing.T) {

	client := &Client{
		ApiKey:    "",
		IsSandbox: true,
	}

	bulkRequest := &BulkRequest{
		Username:      "",
		To:            []string{"+254706496885"},
		Message:       "Hello AT",
		From:          "",
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}

	client.sendBulk(bulkRequest)

}

func TestSendPremium(t *testing.T) {
	premiumRequest := &PremiumRequest{
		Username:      "",
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
