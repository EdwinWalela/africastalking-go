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
	PhoneNumber   string
	Quantity      string
	Unit          string
	Validity      string
	IsPromoBundle string
	MetaData      interface{}
}

/*
DataRequest represents the structure for the request
body to be sent to Africa's Talking api endpoint
*/
type DataRequest struct {
	Username    string
	ProductName string
	Recipients  []Recipient
}

type Entry struct {
	PhoneNumber   string `json:"phoneNumber"`
	Provider      string `json:"provider"`
	Status        string `json:"status"`
	TransactionId string `json:"transactionId"`
	Value         string `json:"value"`
}

type DataResponse struct {
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

func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

func (c *Client) SendMobileData(request *DataRequest) (DataResponse, error) {
	c.Client = &http.Client{}
	url := "https://payments.africastalking.com/mobile/data/request"

	b, _ := json.Marshal(request)

	fmt.Println("REQUEST BODY ", string(b))

	req, _ := http.NewRequest("Post", url, bytes.NewBuffer(b))

	setHeaders(req, c.ApiKey)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error making api request to Africa's Talking API %v", err)
	}

	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)

	var dataResponse DataResponse

	err = json.Unmarshal(responseBody, &dataResponse)

	if err != nil {
		log.Fatalf("Error unmarshalling json", err)
	}

	return DataResponse{
		Entries: dataResponse.Entries,
	}, nil
}
