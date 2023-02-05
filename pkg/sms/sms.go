// Package sms sends Bulk and Premium SMS
package sms

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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

// PremiumRequest represents the request body for the premium SMS request
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

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string // API Key provided by Africa's talking
	IsSandbox bool   // IsSandbox specifies whether to use sandbox or live environment
	client    *http.Client
}

func (c *Client) getUrl() string {
	if c.IsSandbox {
		return "https://api.sandbox.africastalking.com/version1/messaging"
	}
	return "https://api.africastalking.com/version1/messaging"
}

func (c *Client) sendBulk(request *BulkRequest) {
	c.client = &http.Client{}
	data := url.Values{
		"username": {request.Username},
		"to":       {strings.Join(request.To, ",")},
		"message":  {request.Message},
	}
	url := c.getUrl()
	resp, err := c.client.PostForm(url, data)
	if err != nil {
		log.Fatal(err)
	}
	// res := make(map[string]interface{})
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := string(bodyBytes)

	fmt.Println(res)
}
