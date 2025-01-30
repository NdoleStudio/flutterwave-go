package flutterwave

import (
	"context"
	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTransfersService_Rate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.TransferRateResponse()))
	client := New(WithBaseURL(server.URL))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.TransferRateResponse(), *response.Body)
	assert.Equal(t, 624240, rate.Data.Source.Amount)

	// Teardown
	server.Close()
}

func TestTransfersService_Rate_Failure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, `{"error": "internal server error"}`)
	client := New(WithBaseURL(server.URL))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.NotNil(t, err) // Expect an error
	assert.Nil(t, rate)   // The rate should be nil due to failure
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)
	assert.Contains(t, err.Error(), "500") // Ensure error message contains 500

	// Teardown
	server.Close()
}
