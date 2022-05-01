package flutterwave

import "time"

const (
	// HeaderNameVerifHash is the name of the header used to verify your webhook requests.
	HeaderNameVerifHash = "verif-hash"
)

const (
	eventChargeCompleted = "charge.completed"
)

const (
	statusSuccessful = "successful"
	statusFailed     = "failed"
)

// PaymentEventV3 is the payload for webhook requests after a payment
type PaymentEventV3 struct {
	Event string `json:"event"`
	Data  struct {
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
		Customer          struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			PhoneNumber interface{} `json:"phone_number"`
			Email       string      `json:"email"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"customer"`
		Card struct {
			First6Digits string `json:"first_6digits"`
			Last4Digits  string `json:"last_4digits"`
			Issuer       string `json:"issuer"`
			Country      string `json:"country"`
			Type         string `json:"type"`
			Expiry       string `json:"expiry"`
		} `json:"card"`
	} `json:"data"`
	EventType string `json:"event.type"`
}

// IsSuccessful checks if the payment event is successfull
func (event PaymentEventV3) IsSuccessful() bool {
	return event.Event == eventChargeCompleted && event.Data.Status == statusSuccessful
}

// IsFailed checks if the payment failed
func (event PaymentEventV3) IsFailed() bool {
	return event.Event == eventChargeCompleted && event.Data.Status == statusFailed
}
