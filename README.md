# Africa's Talking Golang SDK
![tests](https://github.com/edwinwalela/africastalking-go/actions/workflows/test.yaml/badge.svg) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/edwinwalela/africastalking-go)](https://goreportcard.com/report/github.com/edwinwalela/africastalking-go)  [![Go Reference](https://pkg.go.dev/badge/badge/github.com/edwinwalela/africastalking-go)](https://pkg.go.dev/github.com/edwinwalela/africastalking-go) 

Unofficial Golang SDK for the Africa's Talking API with no external dependencies

## Install   
```bash
go get github.com/edwinwalela/africastalking-go
```

## Quick Start: Sending an SMS
This example makes use of the `sms` package to send bulk sms
```go
package main

import (
	"fmt"
	"github.com/edwinwalela/africastalking-go/pkg/sms"
)

func main() {
	client := &Client{
		ApiKey:    os.Getenv("AT_API_KEY"),
		Username:  os.Getenv("AT_USERNAME"),
		IsSandbox: true,
	}

	bulkRequest := &BulkRequest{
		To:            []string{"+254700000000"},
		Message:       "Hello AT",
		From:          "",
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}

	response, err := client.SendBulk(bulkRequest)

	fmt.Println(response)
}
```
## Local development

Clone repo

```
git clone https://github.com/EdwinWalela/africastalking-go
```

Export Africa's talking credentials (sandbox)

```bash
export AT_USERNAME=sandbox

export AT_API_KEY=xxxxxxxxxxxxx
```

Run tests

```
make test
```
## Resources

- [Code Examples](./examples/)
- [Africa's Talking API Reference](https://developers.africastalking.com/docs/)

## Features

### SMS
- [x] Sending (Bulk & Premium)
- [ ] Premium Subscriptions
- [ ] Fetch Messages
- [ ] Notifications

### USSD
- [ ] Sessions
- [ ] Notifications

### Airtime
- [ ] Sending
- [ ] Query

### Payments
- [ ] C2B
- [ ] B2C
- [ ] B2B
- [ ] Bank
- [ ] Card
- [ ] Query

### Mobile Data
- [ ] Sending

## Contributing

Any contribution, in the form of a suggestion, bug report or pull request, is well accepted.