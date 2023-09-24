/*
Package data enables the ability to sends mobile data
Send Mobile Data API Reference: https://developers.africastalking.com/docs/data/overview
*/

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Recipient struct {
	PhoneNumber   string      `json:"phone_number"`    // PhoneNumber represents the phone number that will be topped up in international format (e.g +234811222333).
	Quantity      int         `json:"quantity"`        // Quantity refers The amount of data
	Unit          string      `json:"unit"`            // Units represents the unit for the specified data quantity, the format is: MB or GB
	Validity      string      `json:"validity"`        //Validity refers to the period of the data bundle’s validity this can be Day, Week, BiWeek, Month, or Quarterly.
	IsPromoBundle string      `json:"is_promo_bundle"` // This is an optional field that can be either true or false.
	MetaData      interface{} `json:"meta_data"`       // A map of any metadata that you would like us to associate with the request.
}

// Request represents the request body send mobile data request
type Request struct {
	Username    string      `json:"username"`    //Username represents the Africa’s Talking application username.
	ProductName string      `json:"productName"` //ProductName refers to the application product name.
	Recipients  []Recipient `json:"recipients"`  // Recipients represents a list of Recipients
}

// Entry is an individual data transaction result
type Entry struct {
	PhoneNumber   string `json:"phoneNumber"`   //The phone number for this transaction.
	Provider      string `json:"provider"`      // This is the name of the service provider.
	Status        string `json:"status"`        // The status of the request associated to this phone number. This could be Queued
	TransactionId string `json:"transactionId"` // A unique id for the request associated to this transaction.
	Value         string `json:"value"`         // The value of data sent. The format of this string is: (3-digit Currency Code)(space)(Decimal Value) e.g KES 200.00
}

// Response represents the response from Africa's Talking API
type Response struct {
	Entries []Entry `json:"entries"` // Entries is a list with an object containing an Entry each corresponding to an individual data transaction result
}

/*
Client represents the http client for communicating with Africa's
Talking Api
*/
type Client struct {
	ApiKey    string       //Api Key provided by Africa's Talking
	Username  string       // Your Africa's Talking application username
	IsSandbox bool         // Is Sandbox specifies whether to use sandbox or live environment
	Client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// setHeaders configures required headers for the HTTP request to Africa's Talking API
func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
}

func (c *Client) Send(request *Request) (Response, error) {
	c.Client = &http.Client{}
	url := "https://payments.africastalking.com/mobile/data/request"

	//Marshal turns the request struct into a []byte to be fed into the http request
	b, _ := json.Marshal(request)

	fmt.Printf("REQUEST BODY : %v", string(b))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))

	setHeaders(req, c.ApiKey)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error making api request to Africa's Talking API %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("Error reading response body into []byte : %v", err)
	}

	fmt.Printf("RESPONSE BODY: %v", string(responseBody))

	var dataResponse Response

	err = json.Unmarshal(responseBody, &dataResponse)

	if err != nil {
		log.Fatalf("Error unmarshalling json : %v", err)
	}

	return Response{
		Entries: dataResponse.Entries,
	}, nil
}
