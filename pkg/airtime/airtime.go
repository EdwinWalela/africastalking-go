/*
Package airtime sends airtime

Africa's Talking API Reference: https://developers.africastalking.com/docs/airtime/sending
*/
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

const (
	KES = "KES" // Kenya Currency Code
	UGX = "UGX" // Uganda Currency Code
	TZS = "TZS" // Tanzania Currency Code
	NGN = "NGN" // Nigeria  Currency Code
	ETB = "ETB" // Ethiopia Currency Code
	MWK = "MWK" // Malawi Currency Code
	ZAR = "ZAR" // South Africa Currency Code
	ZMW = "ZMW" // Zambia  Currency Code
	RWF = "RWF" // Rwanda Currency Code
	GHS = "GHS" // Ghana Currency Code
	XOF = "XOF" // Senegal, Ivory Coast and Cameroon  Currency Code
)

// Recipient represent the target user to receive airtime
type Recipient struct {
	PhoneNumber string  // The number to be topped up "+254xxxxxxx"
	Amount      float64 // Value of airtime to send
	Currency    string  // Currency code of the amount e.g KES,UGX,TZS,NGN,ETB,MWK,ZMW,ZAR
}

// Request represents the request body for the Africa's talking airtime request
type Request struct {
	Recipients []Recipient // Targets to be topped up with airtime
}

// Transaction represents an individual airtime transaction result
type Transaction struct {
	PhoneNumber  string  // Phone number for this transaction
	Amount       float64 // Value of airtime requested
	Currency     string  // Currency code of the amount e.g KES,UGX,TZS,NGN,ETB,MWK,ZMW,ZAR
	Discount     float64 // Discount applied to the requested airtime amount
	Status       string  // Status of the request associated to this phone number
	RequestId    string  // An identifier for the request to this phone number. Only generated if the status of the request is 'sent'
	ErrorMessage string  // Error message for the request associated with this phone number
}

// Response represents the response from Africa's Talking API
type Response struct {
	NumSent       int           // Number of requests sent to the provider
	TotalAmount   float64       // Total value of airtime sent to the provider
	Currency      string        // Currency code of the total amount e.g KES,UGX,TZS,NGN,ETB,MWK,ZMW,ZAR
	TotalDiscount float64       // Total discount applied on the airtime
	Responses     []Transaction // A list of the airtime transaction results
	ErrorMessage  string        // Error message if the entire request was rejected by the API
}

// Client represents the HTTP client responsible for communicating with Africa's Talking API
type Client struct {
	ApiKey    string       // API Key provided by Africa's talking
	Username  string       // Your Africa's talking application username
	IsSandbox bool         // Specifies whether to use sandbox or live environment
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

// Send triggers Africa's Talking airtime API to send Airtime to the specified recipient(s)
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
