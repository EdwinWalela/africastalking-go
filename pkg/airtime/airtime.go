// Package airtime sends airtime
package airtime

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	sandboxURL = "https://api.sandbox.africastalking.com/version1/airtime/send"
	liveURL    = "https://api.africastalking.com/version1/airtime/send"
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
	Discount     float64 // Discount is the discount applied to the requested airtime amount
	Status       string  // Status of the request associated to this phone number
	RequestId    string  // RequestId is an identifier for the request to this phone number. Only generated if the status of the request is 'sent'
	ErrorMessage string  // ErrorMessage is the error for the request associated with this phone number
}

// Response represents the response from Africa's Talking API
type Response struct {
	NumSent       int           // NumSent is the number of requests sent to the provider
	TotalAmount   float64       // TotalAMount is the total value of airtime sent to the provider
	Currency      string        // Currency is the currency code of the total amount
	TotalDiscount float64       // TotalDiscount is the total discount applied on the airtime
	Responses     []Transaction // Responses is a list of the airtime transaction results
	ErrorMessage  string        // ErrorMessage is the error message if the entire request was rejected by the API
}

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // IsSandbox specifies whether to use sandbox or live environment
	Client    *http.Client // HTTP client for making requests to Africa's Talking API
}

// formatRecipients converts the list of recipient to a JSON string
func formatRecipients(recipients []Recipient) string {
	str := "["
	for i, recipient := range recipients {
		str += fmt.Sprintf(`{"phoneNumber":"%s","amount":"%s %.2f"}`, recipient.PhoneNumber, recipient.Currency, recipient.Amount)
		if i != len(recipients)-1 {
			str += ","
		}
	}
	str += "]"
	return str
}

// getRequestBody generates the request body for the airtime HTTP request to Africa's Talking API
func getRequestBody(request *Request, username string) url.Values {
	recipients := formatRecipients(request.Recipients)
	return url.Values{
		"username":   {username},
		"recipients": {recipients},
	}
}

// setHeaders configures required headers for the HTTP request to Africa's Talking API
func setHeaders(request *http.Request, apiKey string) {
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

// formatResponse maps response from Africa's Talking API to the internal Response type
func formatResponse(response *http.Response) (Response, error) {
	res := make(map[string]interface{})
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&res); err != nil {
		return Response{}, err
	}

	responsesData := res["responses"].([]interface{})

	responses := []Transaction{}

	for _, v := range responsesData {
		data := v.(map[string]interface{})
		currency, amountStr, _ := strings.Cut(data["amount"].(string), " ")
		_, discountStr, _ := strings.Cut(data["discount"].(string), " ")
		amount, _ := strconv.ParseFloat(amountStr, 64)
		discount, _ := strconv.ParseFloat(discountStr, 64)
		response := Transaction{
			PhoneNumber:  data["phoneNumber"].(string),
			ErrorMessage: data["errorMessage"].(string),
			Amount:       amount,
			Currency:     currency,
			Discount:     discount,
			Status:       data["status"].(string),
			RequestId:    data["requestId"].(string),
		}
		responses = append(responses, response)
	}

	currency, totalAmountStr, _ := strings.Cut(res["totalAmount"].(string), " ")
	totalAmount, _ := strconv.ParseFloat(totalAmountStr, 64)

	_, totalDiscountStr, _ := strings.Cut(res["totalDiscount"].(string), " ")
	totalDiscount, _ := strconv.ParseFloat(totalDiscountStr, 64)

	return Response{
		ErrorMessage:  res["errorMessage"].(string),
		NumSent:       int(res["numSent"].(float64)),
		TotalAmount:   totalAmount,
		TotalDiscount: totalDiscount,
		Currency:      currency,
		Responses:     responses,
	}, nil
}

func (c *Client) Send(request *Request) (Response, error) {
	c.Client = &http.Client{}
	data := getRequestBody(request, c.Username)
	url := liveURL
	if c.IsSandbox {
		url = sandboxURL
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data.Encode())))
	setHeaders(req, c.ApiKey)
	resp, err := c.Client.Do(req)
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
