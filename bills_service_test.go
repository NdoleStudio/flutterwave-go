package flutterwave

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/araddon/dateparse"
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
	assert.True(t, data.IsSuccessfull())

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
	assert.True(t, data.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestBillsService_GetStatusVerbose(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsGetStatusVerboseResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.GetStatusVerbose(context.Background(), "9300049404444")

	// Assert
	assert.Nil(t, err)

	transactionDate, err := dateparse.ParseAny("2020-03-11T20:19:21.27Z")
	assert.Nil(t, err)

	assert.Equal(t, &BillsStatusVerboseResponse{
		Status:  "success",
		Message: "Bill status fetch successful",
		Data: struct {
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
		}{
			"NGN",
			"+23490803840303",
			"One Time",
			"500.0000",
			"AIRTIME",
			"9MOBILE",
			10,
			transactionDate,
			"NG",
			"CF-FLYAPI-20200311081921359990",
			nil,
			"FLY-API-NG-AIRTIME-9MOBILE",
			"successful",
		},
	}, data)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.True(t, data.IsSuccessfull())

	// Teardown
	server.Close()
}
