package flutterwave

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// transactionsService is the API client for the `/v3/transactions` endpoint
type transactionsService service

// Verify the final status of a transaction
//
// API Docs: https://developer.flutterwave.com/reference/endpoints/transactions
func (service *transactionsService) Verify(ctx context.Context, transactionID int64) (*TransactionResponse, *Response, error) {
	uri := fmt.Sprintf("/v3/transactions/%d/verify", transactionID)

	request, err := service.client.newRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var data TransactionResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil
}
