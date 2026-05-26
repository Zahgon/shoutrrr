package gotify

import (
	"net/http"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/containrrr/shoutrrr/pkg/util/jsonclient"
)

// Service providing Gotify as a notification service
type Service struct {
	standard.Standard
	config     *Config
	pkr        format.PropKeyResolver
	httpClient *http.Client
	client     jsonclient.Client
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// If DisableTLS is specified, we might still need to disable TLS verification
// since the default configuration of Gotify redirects HTTP to HTTPS
// Note that this cannot be overridden using params, only using the config URL

// Set a reasonable timeout to prevent one bad transfer from block all subsequent ones

const tokenChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_"

// The validation rules have been taken directly from the Gotify source code.
// These will have to be adapted in case of a change:
// https://github.com/gotify/server/blob/ad157a138b4985086c484a7aabfc2deada5a33dd/auth/token.go#L8
func isTokenValid(token string) bool { _ = "STUB: not implemented"; return false }

func buildURL(config *Config) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Send a notification message to Gotify
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHTTPClient is only supposed to be used for mocking the httpclient when testing
func (service *Service) GetHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }
