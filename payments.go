package flutterwave

// GetPaymentLinkRequest is data needed to create a payment link
type GetPaymentLinkRequest struct {
	TransactionReference string                       `json:"tx_ref"`
	Amount               string                       `json:"amount"`
	Currency             string                       `json:"currency"`
	Meta                 map[string]interface{}       `json:"meta"`
	RedirectURL          string                       `json:"redirect_url"`
	Customer             GetPaymentLinkCustomer       `json:"customer"`
	Customizations       GetPaymentLinkCustomizations `json:"customizations"`
}

// GetPaymentLinkCustomer contains the customer details.
type GetPaymentLinkCustomer struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
}

// GetPaymentLinkCustomizations contains options to customize the look of the payment modal.
type GetPaymentLinkCustomizations struct {
	Title string `json:"title"`
	Logo  string `json:"logo"`
}

// GetPaymentLinkResponse is the data returned after creating a payment link
type GetPaymentLinkResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Link string `json:"link"`
	} `json:"data"`
}
