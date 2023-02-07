/*
Package voice provides Africa's Talking Voice services.

Africa's Talking API Reference: https://developers.africastalking.com/docs/voice/overview
*/
package voice

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	callLiveURL        = "https://voice.africastalking.com/call"
	callSandboxURL     = "https://voice.sandbox.africastalking.com/call"
	transferLiveUrl    = "https://voice.africastalking.com/callTransfer"
	transferSandboxURL = "https://voice.sandbox.africastalking.com/callTransfer"
)

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// CallRequest represents the body to be sent to the Voice API when initiating a call
type CallRequest struct {
	From            string   // Your Africa's Talking phone number "+254xxxxxxxx"
	To              []string // Recipients' phone numbers
	ClientRequestId string   // Identifier sent to the registered Callback URL that can be used to tag the call
}

// CallTransferRequest represents the body to be sent to the Voice API when transferring a call
type CallTransferRequest struct {
	SessionId    string // Identifier of the on going call (Required)
	PhoneNumber  string // Phone number to transfer the call to (Required)
	CallLeg      string // Call leg to transfer the call to. Either 'caller' or 'callee' Default 'callee' (Optional)
	HoldMusicUrl string // URL of the media file to be played when user is on hold. 'http://xxxxxx'
}

// Recipient represents the status of the call of an individual user specified in Request
type Recipient struct {
	PhoneNumber string // Recipient's phone number
	Status      string // Status of the request:e.g "Queued","InvalidPhoneNumber","DestinationNotSupported","Insufficient Credit"
	SessionId   string // A unique identifier for the request associated to this phone number
}

// CallResponse represents the response after initiating a Call
type CallResponse struct {
	Recipients   []Recipient // List of recipients and their status
	ErrorMessage string      // Error message if the entire request was rejected by the API
}

// CallTransferResponse represents the
type CallTransferResponse struct {
	Status       string // Status of the call transfer request. 'Success' or 'Aborted'
	ErrorMessage string // Reason why the transfer was aborted
}

// setHeaders configures required headers for the HTTP request to Africa's Talking API
func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

// getRequestBody generates the request body for the voice HTTP request to Africa's Talking API
func getCallRequestBody(request *CallRequest, username string) url.Values {
	return url.Values{
		"username":        {username},
		"from":            {request.From},
		"to":              {strings.Join(request.To, ",")},
		"clientRequestId": {request.ClientRequestId},
	}
}

// getCallTransferRequestBody generates the request body for the call transfer HTTP request to Africa's Talking API
func getCallTransferRequestBody(request *CallTransferRequest, username string) url.Values {
	return url.Values{
		"username":     {username},
		"sessionId":    {request.SessionId},
		"callLeg":      {request.CallLeg},
		"holdMusicUrl": {request.HoldMusicUrl},
	}
}

// formatCallResponse maps response from Africa's Talking call API to the internal Response type
func formatCallResponse(response *http.Response) (CallResponse, error) {
	res := make(map[string]interface{})
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&res); err != nil {
		return CallResponse{}, err
	}

	entries := res["entries"].([]interface{})
	recipients := []Recipient{}

	for _, entry := range entries {
		data := entry.(map[string]interface{})
		recipient := Recipient{
			PhoneNumber: data["phoneNumber"].(string),
			Status:      data["status"].(string),
			SessionId:   data["sessionId"].(string),
		}
		recipients = append(recipients, recipient)
	}

	return CallResponse{
		ErrorMessage: res["errorMessage"].(string),
		Recipients:   recipients,
	}, nil
}

// formatCallTransferResponse maps response from Africa's Talking call transfer API to the internal Response type
func formatCallTransferResponse(response *http.Response) (CallTransferResponse, error) {
	res := make(map[string]string)
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&res); err != nil {
		return CallTransferResponse{}, err
	}
	return CallTransferResponse{
		Status:       res["status"],
		ErrorMessage: res["errorMessage"],
	}, nil
}

/*
Call makes an outbound call through Africa's Talking Voice API.

When the call is picked, Africa's Talking will call your callback url.

API Reference: https://developers.africastalking.com/docs/voice/handle_calls
*/
func (c *Client) Call(request *CallRequest) (CallResponse, error) {
	c.client = &http.Client{}
	data := getCallRequestBody(request, c.Username)
	url := callLiveURL
	if c.IsSandbox {
		url = callSandboxURL
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	setHeaders(req, c.ApiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return CallResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return CallResponse{}, err
		}
		res := string(bodyBytes)
		return CallResponse{}, errors.New(res)
	}

	return formatCallResponse(resp)
}

// Transfer transfers a call to another number.  Only works in live environment
func (c *Client) Transfer(request *CallTransferRequest) (CallTransferResponse, error) {
	c.client = &http.Client{}
	data := getCallTransferRequestBody(request, c.Username)
	url := transferLiveUrl
	if c.IsSandbox {
		url = transferSandboxURL
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	setHeaders(req, c.ApiKey)
	resp, err := c.client.Do(req)

	if err != nil {
		return CallTransferResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return CallTransferResponse{}, err
		}
		res := string(bodyBytes)
		return CallTransferResponse{}, errors.New(res)
	}

	return formatCallTransferResponse(resp)
}
