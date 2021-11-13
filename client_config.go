package flutterwave

import (
	"net/http"
)

type clientConfig struct {
	httpClient *http.Client
	secretKey  string
	baseURL    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		secretKey:  "",
		baseURL:    "https://api.flutterwave.com",
	}
}
