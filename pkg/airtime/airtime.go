// Package airtime sends airtime
package airtime

// Recipient represent the target user to receive airtime
type Recipient struct {
	PhoneNumber string  // Phonenumber is the number to be topped up "+254xxxxxxx"
	Amount      float64 // Amount is the value of airtime to send
	Currency    string  // Currency is the currency code of the amount e.g KES
}

// Request represents the request body for the Africa's talking airtime request
type Request struct {
	Username   string      // Username is your africa's talking application username
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
