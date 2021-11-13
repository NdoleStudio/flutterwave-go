package flutterwave

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBillsService_CreatePayment(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsCreateDStvPaymentResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.CreatePayment(context.Background(), &BillsCreatePaymentRequest{
		Country:    "NG",
		Customer:   "7034504232",
		Amount:     100,
		Recurrence: "ONCE",
		Type:       "DSTV",
		Reference:  uuid.New().String(),
		BillerName: "DSTV",
	})

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, &BillsCreatePaymentResponse{
		Status:  "success",
		Message: "Bill payment successful",
		Data: struct {
			PhoneNumber string `json:"phone_number"`
			Amount      int    `json:"amount"`
			Network     string `json:"network"`
			FlwRef      string `json:"flw_ref"`
			TxRef       string `json:"tx_ref"`
		}{
			"+23490803840303",
			500,
			"9MOBILE",
			"CF-FLYAPI-20200311081921359990",
			"BPUSSD1583957963415840",
		},
	}, data)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}

func TestBillsService_Validate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsValidateDstvResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.Validate(context.Background(), "CB177", "BIL099", "08038291822")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, &BillsValidateResponse{
		Status:  "success",
		Message: "Item validated successfully",
		Data: struct {
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
		}{
			"00",
			nil,
			"Successful",
			"MTN",
			"BIL099",
			"08038291822",
			"AT099",
			nil,
			100,
			0,
			0,
		},
	}, data)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
