package flutterwave

import (
	"context"
	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTransfersService_Query(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.TransferRateResponse()))
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Transfers.Query(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.TransferRateResponse(), *response.Body)

	// Teardown
	server.Close()
}