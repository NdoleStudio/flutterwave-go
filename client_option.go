package flutterwave

import (
	"net/http"
	"strings"
)

// ClientOption are options for constructing a client
type ClientOption interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the flutterwave API
func WithBaseURL(baseURL string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithSecretKey set's the secret key used to authorize requests to the flutterwave API
// See: https://developer.flutterwave.com/docs/api-keys
func WithSecretKey(secretKey string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.secretKey = secretKey
	})
}
