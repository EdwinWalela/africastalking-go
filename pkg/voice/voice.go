/*
Package voice provides Africa's Talking Voice services.

Africa's Talking API Reference: https://developers.africastalking.com/docs/voice/overview
*/
package voice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	liveURL    = "https://voice.africastalking.com/call"
	sandboxURL = "https://voice.sandbox.africastalking.com/call"
)

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// Request represents the body to be sent to the Voice API
type Request struct {
	From            string   // Your Africa's Talking phone number "+254xxxxxxxx"
	To              []string // Recipients' phone numbers
	ClientRequestId string   // Identifier sent to the registered Callback URL that can be used to tag the call
}

// Recipient represents the status of the call of an individual user specified in Request
type Recipient struct {
	PhoneNumber string // Recipient's phone number
	Status      string // Status of the request:e.g "Queued","InvalidPhoneNumber","DestinationNotSupported","Insufficient Credit"
	SessionId   string // A unique identifier for the request associated to this phone number
}

// Response represents the response from Africa's Talking Voice API
type Response struct {
	Recipients   []Recipient // List of recipients and their status
	ErrorMessage string      // Error message if the entire request was rejected by the API
}

// setHeaders configures required headers for the HTTP request to Africa's Talking API
func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

// getRequestBody generates the request body for the voice HTTP request to Africa's Talking API
func getRequestBody(request *Request, username string) url.Values {
	return url.Values{
		"username":        {username},
		"from":            {request.From},
		"to":              {strings.Join(request.To, ",")},
		"clientRequestId": {request.ClientRequestId},
	}
}

// formatResponse maps response from Africa's Talking API to the internal Response type
func formatResponse(response *http.Response) (Response, error) {
	res := make(map[string]interface{})
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&res); err != nil {
		return Response{}, err
	}

	fmt.Println(res)
	return Response{}, nil
}

/*
Call makes an outbound call through Africa's Talking Voice API.

When the call is picked, Africa's Talking will call your callback url.

API Reference: https://developers.africastalking.com/docs/voice/handle_calls
*/
func (c *Client) Call(request *Request) (Response, error) {
	c.client = &http.Client{}
	data := getRequestBody(request, c.Username)
	url := liveURL
	if c.IsSandbox {
		url = sandboxURL
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	setHeaders(req, c.ApiKey)
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
