package jsonclient

import (
	"net/http"
)

// ContentType is the default mime type for JSON
const ContentType = "application/json"

// DefaultClient is the singleton instance of jsonclient using http.DefaultClient
var DefaultClient = NewClient()

// Get fetches url using GET and unmarshals into the passed response using DefaultClient
func Get(url string, response interface{}) error { _ = "STUB: not implemented"; return nil }

// Post sends request as JSON and unmarshals the response JSON into the supplied struct using DefaultClient
func Post(url string, request interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Client is a JSON wrapper around http.Client
type client struct {
	httpClient *http.Client
	headers    http.Header
	indent     string
}

// NewClient returns a new JSON Client using the default http.Client
func NewClient() Client { _ = "STUB: not implemented"; return *new(Client) }

// NewWithHTTPClient returns a new JSON Client using the specified http.Client
func NewWithHTTPClient(httpClient *http.Client) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// Headers return the default headers for requests
func (c *client) Headers() http.Header {
	_ = "STUB: not implemented"

	// Get fetches url using GET and unmarshals into the passed response
	return *new(http.Header)
}

func (c *client) Get(url string, response interface{}) error { _ = "STUB: not implemented"; return nil }

// Post sends request as JSON and unmarshals the response JSON into the supplied struct
func (c *client) Post(url string, request interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// If the request is a string, just pass it through without serializing

func (c *client) ErrorResponse(err error, response interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func parseResponse(res *http.Response, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
