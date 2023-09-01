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
	PhoneNumber   string      `json:"phone_number"`
	Quantity      int         `json:"quantity"`
	Unit          string      `json:"unit"`
	Validity      string      `json:"validity"`
	IsPromoBundle string      `json:"is_promo_bundle"`
	MetaData      interface{} `json:"meta_data"`
}

// Request represents the request body send mobile data request
type Request struct {
	Username    string      `json:"username"`
	ProductName string      `json:"productName"`
	Recipients  []Recipient `json:"recipients"`
}

// Entry is an individual data transaction result
type Entry struct {
	PhoneNumber   string `json:"phoneNumber"`
	Provider      string `json:"provider"`
	Status        string `json:"status"`
	TransactionId string `json:"transactionId"`
	Value         string `json:"value"`
}

// Response represents the response from Africa's Talking API
type Response struct {
	Entries []Entry `json:"entries"`
}

/*
Client represents the http client for communicating with Africa's
Talking Api
*/
type Client struct {
	ApiKey    string       //Api Key provided by Africa's Talking
	Username  string       // Your Africa's Talking application username
	IsSandbox bool         // Is Sandbox specifies whether to use sandbox or live environment
	Client    *http.Client //
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

	// fmt.Printf("REQUEST BODY : %v", string(b))

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
