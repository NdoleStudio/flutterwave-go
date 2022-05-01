package flutterwave

import (
	"context"
	"encoding/json"
	"net/http"
)

// billsService is the API client for the `/gateway` endpoint
type billsService service

// CreatePayment creates bill payments.
//
// API Docs: https://developer.flutterwave.com/reference/create-a-bill-payment
func (service *billsService) CreatePayment(ctx context.Context, payload *BillsCreatePaymentRequest) (*BillsCreatePaymentResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/v3/bills", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var data BillsCreatePaymentResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil
}

// Validate validates services like DStv smartcard number, Meter number etc.
//
// API Docs: https://developer.flutterwave.com/reference/validate-bill-service
func (service *billsService) Validate(ctx context.Context, itemCode string, billerCode string, customer string) (*BillsValidateResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v3/bill-items/"+itemCode+"/validate", nil)
	if err != nil {
		return nil, nil, err
	}

	request = service.client.addURLParams(request, map[string]string{
		"code":     billerCode,
		"customer": customer,
	})

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var data BillsValidateResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil
}

// GetStatusVerbose gets the verbose status of a bill payment.
//
// API Docs: https://developer.flutterwave.com/reference/get-status-of-a-bill-payment
func (service *billsService) GetStatusVerbose(ctx context.Context, transactionReference string) (*BillsStatusVerboseResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/v3/bills/"+transactionReference, map[string]int{"verbose": 1})
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var data BillsStatusVerboseResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, err
	}

	return &data, response, nil
}
