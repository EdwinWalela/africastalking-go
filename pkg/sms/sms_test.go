package sms

import (
	"testing"
	"time"
)

func TestSendBulk(t *testing.T) {
	bulkRequest := &BulkRequest{
		Username:      "",
		To:            []string{""},
		Message:       "Hello AT",
		From:          "",
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}
	_ = bulkRequest
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
