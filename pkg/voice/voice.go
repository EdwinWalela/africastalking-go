/*
Package voice provides Africa's Talking Voice services.

Africa's Talking API Reference: https://developers.africastalking.com/docs/voice/overview
*/
package voice

import "net/http"

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	Client    *http.Client // HTTP client for making requests to Africa's Talking API
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
