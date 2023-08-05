package flutterwave

import (
	"context"
	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTransactionsService_Refund(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.TransactionRefundResponse()))
	client := New(WithBaseURL(server.URL))

	// Act
	refund, response, err := client.Transactions.Refund(context.Background(), 123, 200)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.TransactionRefundResponse(), *response.Body)
	assert.Equal(t, 75923, refund.Data.ID)

	// Teardown
	server.Close()
}

func TestTransactionsService_RefundWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, "")
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Transactions.Refund(context.Background(), 123, 200)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
