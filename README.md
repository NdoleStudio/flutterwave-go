# flutterwave-go

[![Build](https://github.com/NdoleStudio/flutterwave-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/flutterwave-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/flutterwave-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/flutterwave-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/flutterwave-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/flutterwave-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/flutterwave-go)](https://goreportcard.com/report/github.com/NdoleStudio/flutterwave-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/flutterwave-go)](https://github.com/NdoleStudio/flutterwave-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/flutterwave-go?color=brightgreen)](https://github.com/NdoleStudio/flutterwave-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/flutterwave-go)](https://pkg.go.dev/github.com/NdoleStudio/flutterwave-go)


This package provides a `go` client for interacting with the [Flutterwave API](https://developer.flutterwave.com/docs)

## Installation

`flutterwave-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/flutterwave-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/flutterwave-go"
```

## Implemented

- **[Bills](#bills)**
  - `POST /bills/`: Create a bill payment
  - `GET /bill-items/{item_code}/validate`: Validate services like DStv smartcard number, Meter number etc.
  - `GET /bills/{reference}`: Get the verbose status of a bill payment
- **Payments**
  - `GET /v3/transactions/:id/verify`: Verify a transaction
  - `POST /v3/transactions/:id/refund`: Create a Refund

## Usage

### Initializing the Client

An instance of the `flutterwave` client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/flutterwave-go"
)

func main()  {
	client := flutterwave.New(
		flutterwave.WithSecretKey("" /* flutterwave Secret Key */),
	)
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
data, httpResponse, err := flutterwaveClient.Bills.Create(context.Background(), request)
if err != nil {
    //handle error
}
```

### BILLS

#### Create a bill payment

`POST /bills/`: Create a bill payment

```go
response, _, err := flutterwaveClient.Bills.CreatePayment(context.Background(), &BillsCreatePaymentRequest{
    Country:    "NG",
    Customer:   "7034504232",
    Amount:     100,
    Recurrence: "ONCE",
    Type:       "DSTV",
    Reference:  uuid.New().String(),
    BillerName: "DSTV",
})

if err != nil {
    log.Fatal(err)
}

log.Println(response.Status) // success
```


#### Create a bill payment

`GET /bill-items/{item_code}/validate`: validate services like DSTV smartcard number, Meter number etc.

```go
response, _, err := flutterwaveClient.Bills.Validate(context.Background(), "CB177", "BIL099", "08038291822")

if err != nil {
    log.Fatal(err)
}

log.Println(response.Status) // success
```

#### Get verbose status of a bill payment

`GET /bills/{reference}`: get the verbose status of a bill purchase

```go
response, _, err := flutterwaveClient.Bills.GetStatusVerbose(context.Background(), "9300049404444")

if err != nil {
    log.Fatal(err)
}

log.Println(response.Status) // success
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
