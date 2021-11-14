package flutterwave

import "time"

// BillsCreatePaymentRequest is data needed to creat a payment
type BillsCreatePaymentRequest struct {
	Country    string `json:"country"`
	Customer   string `json:"customer"`
	Amount     int    `json:"amount"`
	Recurrence string `json:"recurrence,omitempty"`
	Type       string `json:"type"`
	Reference  string `json:"reference,omitempty"`
	BillerName string `json:"biller_name,omitempty"`
}

// BillsCreatePaymentResponse is the data returned after creating a payment
type BillsCreatePaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		PhoneNumber string `json:"phone_number"`
		Amount      int    `json:"amount"`
		Network     string `json:"network"`
		FlwRef      string `json:"flw_ref"`
		TxRef       string `json:"tx_ref"`
	} `json:"data"`
}

// IsSuccessfull determines if the bill payment was successfull
func (response BillsCreatePaymentResponse) IsSuccessfull() bool {
	return response.Status == "success" && response.Data.TxRef != ""
}

// BillsValidateResponse is the response after validating a bill service
type BillsValidateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ResponseCode    string      `json:"response_code"`
		Address         interface{} `json:"address"`
		ResponseMessage string      `json:"response_message"`
		Name            string      `json:"name"`
		BillerCode      string      `json:"biller_code"`
		Customer        string      `json:"customer"`
		ProductCode     string      `json:"product_code"`
		Email           interface{} `json:"email"`
		Fee             int         `json:"fee"`
		Maximum         int         `json:"maximum"`
		Minimum         int         `json:"minimum"`
	} `json:"data"`
}

// IsSuccessfull determines if the bill validation was successfull
func (response BillsValidateResponse) IsSuccessfull() bool {
	return response.Status == "success" && response.Data.Customer != ""
}

// BillsStatusVerboseResponse is the verbose response of a bill payment
type BillsStatusVerboseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Currency        string      `json:"currency"`
		CustomerID      string      `json:"customer_id"`
		Frequency       string      `json:"frequency"`
		Amount          string      `json:"amount"`
		Product         string      `json:"product"`
		ProductName     string      `json:"product_name"`
		Commission      int         `json:"commission"`
		TransactionDate time.Time   `json:"transaction_date"`
		Country         string      `json:"country"`
		TxRef           string      `json:"tx_ref"`
		Extra           interface{} `json:"extra"`
		ProductDetails  string      `json:"product_details"`
		Status          string      `json:"status"`
	} `json:"data"`
}

// IsSuccessfull determines if the transaction was successfull
func (response BillsStatusVerboseResponse) IsSuccessfull() bool {
	return response.Data.Status == "successful" && response.Data.TxRef != ""
}
