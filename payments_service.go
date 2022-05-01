package flutterwave

import (
	"context"
	"encoding/json"
	"net/http"
)

// paymentsService is the API client for the `/gateway` endpoint
type paymentsService service

// GetPaymentLink Call flutterwave to get a payment link, redirect your customer to the link, and flutterwave will redirect back when payment is done.
//
// API Docs: https://developer.flutterwave.com/docs/collecting-payments/standard
func (service *paymentsService) GetPaymentLink(ctx context.Context, payload *GetPaymentLinkRequest) (*GetPaymentLinkResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v3/payments", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var data GetPaymentLinkResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil
}
