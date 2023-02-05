// Package sms sends Bulk and Premium SMS
package sms

import (
	"fmt"
	"time"
)

// BulkRequest represents the request body for the bulk SMS request
type BulkRequest struct {
	Username      string        // Username is your africa's talking application username (required)
	To            []string      // To specifies recipients` phone numbers "+2547XXXXXXXX"` (required)
	Message       string        // Message is the contents of the sms to be sent (required)
	From          string        // From is your registered short code or alphanumerics, defaults to AFRICASTKNG (optional)
	BulkSMSMode   bool          // BulkSMSMode determines who gets billed for a message sent, default value is true (sender is billed), must be true for bulk messages (optional)
	Enqueue       bool          // If enabled, the API will store the messages in a queue and send them out asynchronously after responding to the request (optional)
	RetryDuration time.Duration // RetryDuration specifies the number of hours your subscription message should be retried in case it’s not delivered to the subscriber (optional)
}

type PremiumRequest struct {
	Username      string        // Username your africa's talking application username (required)
	To            []string      // To specifies recipients` phone numbers "+2547XXXXXXXX"` (required)
	Message       string        // Message is the contents of the sms to be sent (required)
	From          string        // From is your registered short code or alphanumerics, defaults to AFRICASTKNG (optional)
	Keyword       string        // Keyword to be used for a premium service (optional)
	Enqueue       bool          // If enabled, the API will store the messages in a queue and send them out asynchronously after responding to the request (optional)
	LinkId        string        // LinkId is used to send OnDemand messages. The linkId is forwarded to your application when the user sends a message to your service. (optional)
	RetryDuration time.Duration // RetryDuration specifies the number of hours your subscription message should be retried in case it’s not delivered to the subscriber (optional)
	RequestId     string        // RequestId is a client specified request identifier. Returned as part of the http dlr callback
}

func SendBulk() {
	fmt.Println("Sending bulk")
}

func SendPremium() {
	fmt.Println("Sending premium")
}
