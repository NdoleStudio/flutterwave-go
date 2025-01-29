package flutterwave

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// transfersService is the API client for the `/v3/transfers` endpoint
type transfersService service

// Estimate the Transfer Rate of a transaction
//
// API Docs: https://developer.flutterwave.com/reference/check-transfer-rates
func (service *transfersService) Query(ctx context.Context, amount int, destination_currency, source_currency string) (*TransferRateResponse, *Response, error) {
	uri := "/v3/transfers/rates"

	request, err := service.client.newRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	// Adding the parameters
	requestWithParams := service.client.addURLParams(request, map[string]string{
		"amount": strconv.Itoa(amount),
		"destination_currency": destination_currency,
		"source_currency": source_currency,
	})

	response, err := service.client.do(requestWithParams)
	if err != nil {
		return nil, response, err
	}

	var data TransferRateResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil

}
