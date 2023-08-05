package flutterwave

import "time"

// TransactionResponse is data returned when querying a transaction
type TransactionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID                int64     `json:"id"`
		TxRef             string    `json:"tx_ref"`
		FlwRef            string    `json:"flw_ref"`
		DeviceFingerprint string    `json:"device_fingerprint"`
		Amount            int       `json:"amount"`
		Currency          string    `json:"currency"`
		ChargedAmount     int       `json:"charged_amount"`
		AppFee            float64   `json:"app_fee"`
		MerchantFee       int       `json:"merchant_fee"`
		ProcessorResponse string    `json:"processor_response"`
		AuthModel         string    `json:"auth_model"`
		IP                string    `json:"ip"`
		Narration         string    `json:"narration"`
		Status            string    `json:"status"`
		PaymentType       string    `json:"payment_type"`
		CreatedAt         time.Time `json:"created_at"`
		AccountID         int       `json:"account_id"`
		Card              struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Token        string `json:"token"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
		Meta          interface{} `json:"meta"`
		AmountSettled float64     `json:"amount_settled"`
		Customer      struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			PhoneNumber string    `json:"phone_number"`
			Email       string    `json:"email"`
			CreatedAt   time.Time `json:"created_at"`
		} `json:"customer"`
	} `json:"data"`
}

// RefundTransactionResponse is the payload generated when a transaction is refunded
type RefundTransactionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID             int    `json:"id"`
		AccountID      int    `json:"account_id"`
		TxID           int    `json:"tx_id"`
		FlwRef         string `json:"flw_ref"`
		WalletID       int    `json:"wallet_id"`
		AmountRefunded int    `json:"amount_refunded"`
		Status         string `json:"status"`
		Destination    string `json:"destination"`
		Meta           struct {
			Source string `json:"source"`
		} `json:"meta"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"data"`
}
