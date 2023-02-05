// Package sms sends Bulk and Premium SMS
package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	To            []string      // To specifies recipients` phone numbers "+2547XXXXXXXX"` (required)
	Message       string        // Message is the contents of the sms to be sent (required)
	From          string        // From is your registered short code or alphanumerics, defaults to AFRICASTKNG (optional)
	Keyword       string        // Keyword to be used for a premium service (optional)
	Enqueue       bool          // If enabled, the API will store the messages in a queue and send them out asynchronously after responding to the request (optional)
	LinkId        string        // LinkId is used to send OnDemand messages. The linkId is forwarded to your application when the user sends a message to your service. (optional)
	RetryDuration time.Duration // RetryDuration specifies the number of hours your subscription message should be retried in case it’s not delivered to the subscriber (optional)
	RequestId     string        // RequestId is a client specified request identifier. Returned as part of the http dlr callback
}

// Recipient represents a recipient who was included in the original request
type Recipient struct {
	Status     string // Status indicates whether the SMS was sent to the recipient or not
	StatusCode uint16 // StatusCode is the status of the request
	Number     string // Number is the recipient's phone number
	Cost       string // Cost is the amount incurred to send this SMS
	MessageId  string // MessageId received when the sms was sent
}

// Response represents the response from Africa's Talking API
type Response struct {
	Message    string // Message is the summary of the total number of recipients the sms was sent to and total cost
	Recipients []Recipient
}

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	apiKey    string       // API Key provided by Africa's talking
	username  string       // Your Africa's talking application username
	isSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// getUrl returns Africa's Talking SMS URL based on the specified client environment
func getUrl(isSandbox bool) string {
	if isSandbox {
		return "https://api.sandbox.africastalking.com/version1/messaging"
	}
	return "https://api.africastalking.com/version1/messaging"
}

// getRequestBody generates the request body for the bulk SMS HTTP request to Africa's Talking API
func getBulkRequestBody(request *BulkRequest, username string, isSandbox bool) url.Values {
	data := url.Values{
		"username": {username},
		"to":       {strings.Join(request.To, ",")},
		"message":  {request.Message},
	}
	if !isSandbox {
		data.Add("from", request.From)
	}
	if request.Enqueue {
		data.Set("enqueue", "1")
	} else {
		data.Set("enqueue", "0")
	}
	if request.BulkSMSMode {
		data.Set("bulkSMSMode", "1")
	} else {
		data.Set("bulkSMSMode", "0")
	}
	data.Set("retryDurationInHours", fmt.Sprintf("%.0f", request.RetryDuration.Abs().Hours()))
	return data
}

// setHeaders configures required headers for the HTTP request to Africa's Talking API
func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

// formartResponse maps response from Africa's Talking API to the internal Response type
func formatResponse(response *http.Response) (Response, error) {
	res := make(map[string]interface{})
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&res); err != nil {
		return Response{}, err
	}

	data := res["SMSMessageData"].(map[string]interface{})
	recepientsData := data["Recipients"].([]interface{})
	recipients := []Recipient{}

	for _, v := range recepientsData {
		data := v.(map[string]interface{})
		recipient := Recipient{
			StatusCode: uint16(data["statusCode"].(float64)),
			Number:     data["number"].(string),
			Cost:       data["cost"].(string),
			MessageId:  data["messageId"].(string),
			Status:     data["status"].(string),
		}
		recipients = append(recipients, recipient)
	}

	return Response{
		Message:    data["Message"].(string),
		Recipients: recipients,
	}, nil
}

// sendBulk sends Bulk SMS using the Africa's Talking API
func (c *Client) SendBulk(request *BulkRequest) (Response, error) {
	c.client = &http.Client{}
	data := getBulkRequestBody(request, c.username, c.isSandbox)
	url := getUrl(c.isSandbox)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	setHeaders(req, c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return Response{}, err
		}
		res := string(bodyBytes)
		return Response{}, errors.New(res)
	}

	return formatResponse(resp)
}
