package testutils

import (
	"net/url"

	"github.com/jarcoal/httpmock"
)

// URLMust creates a url.URL from the given rawURL and fails the test if it cannot be parsed
func URLMust(rawURL string) *url.URL { _ = "STUB: not implemented"; return nil }

// JSONRespondMust creates a httpmock.Responder with the given response as the body, and fails the test if it cannot be created
func JSONRespondMust(code int, response interface{}) httpmock.Responder {
	_ = "STUB: not implemented"
	return *new(httpmock.Responder)
}
