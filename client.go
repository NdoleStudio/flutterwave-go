package flutterwave

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Pre-defined error messages to be shared across services.
var (
	ErrCouldNotConstructNewRequest = errors.New("could not construct new request")
	ErrRequestFailure  = errors.New("request failed")
	ErrUnmarshalFailure = errors.New("failed to unmarshal response")
)

type service struct {
	client *Client
}

// Client is the flutterwave API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient *http.Client
	common     service
	secretKey  string
	baseURL    string

	Bills        *billsService
	Payments     *paymentsService
	Transactions *transactionsService
	Transfers *transfersService
}

// New creates and returns a new flutterwave.Client from a slice of flutterwave.ClientOption.
func New(options ...ClientOption) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient: config.httpClient,
		secretKey:  config.secretKey,
		baseURL:    config.baseURL,
	}

	client.common.client = client
	client.Bills = (*billsService)(&client.common)
	client.Payments = (*paymentsService)(&client.common)
	client.Transactions = (*transactionsService)(&client.common)
	client.Transfers = (*transfersService)(&client.common)
	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.secretKey)

	return req, nil
}

// addURLParams adds urls parameters to an *http.Request
func (client *Client) addURLParams(request *http.Request, params map[string]string) *http.Request {
	q := request.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	return request
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(io.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := io.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
