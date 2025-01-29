package flutterwave

// ExchangeRateResponse is data returned when querying a transaction
type TransferRateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Rate        float64 `json:"rate"`
		Source      Payment `json:"source"`
		Destination Payment `json:"destination"`
	}
}

type Payment struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}
