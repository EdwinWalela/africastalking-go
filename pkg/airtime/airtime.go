// Package airtime sends airtime
package airtime

import (
	"fmt"
	"net/http"
	"net/url"
)

// Recipient represent the target user to receive airtime
type Recipient struct {
	PhoneNumber string  // Phonenumber is the number to be topped up "+254xxxxxxx"
	Amount      float64 // Amount is the value of airtime to send
	Currency    string  // Currency is the currency code of the amount e.g KES
}

// Request represents the request body for the Africa's talking airtime request
type Request struct {
	Recipients []Recipient // Recipients are the targets to be topped up
}

// Transaction represents an individual airtime transaction result
type Transaction struct {
	PhoneNumber  string  // Phone number for this transaction
	Amount       float64 // Amount is the value of airtime requested
	Currency     string  // Currency is the currency code of the amount
	Discount     string  // Discount is the discount applied to the requested airtime amount
	Status       string  // Status of the request associated to this phone number
	RequestId    string  // RequestId is an identifier for the request to this phone number. Only generated if the status of the request is 'sent'
	ErrorMessage string  // ErrorMessage is the error for the request associated with this phone number
}

// Response represents the response from Africa's Talking API
type Response struct {
	NumSent       int     // NumSent is the number of requests sent to the provider
	TotalAmount   float64 // TotalAMount is the total value of airtime sent to the provider
	Currency      string  // Currency is the currency code of the total amount
	TotalDiscount float64 // TotalDiscount is the total discount applied on the airtime
	ErrorMessage  string  // ErrorMessage is the error message if the entire request was rejected by the API
}

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	Client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// getRequestBody generates the request body for the airtime HTTP request to Africa's Talking API
func getRequestBody(request *Request, username string) url.Values {
	data := url.Values{
		"username": {username},
	}
	recipients := "["
	for i, recipient := range request.Recipients {
		recipients += fmt.Sprintf(`{"phoneNumber":"%s"."amount":"%s %.2f"}`, recipient.PhoneNumber, recipient.Currency, recipient.Amount)
		if i != len(request.Recipients)-1 {
			recipients += ","
		}
	}
	recipients += "]"
	return data
}

func (c *Client) Send(request *Request) (Response, error) {
	c.Client = &http.Client{}
	body := getRequestBody(request, c.Username)
	_ = body
	return Response{}, nil
}
