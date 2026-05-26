// Package ntfy implements Ntfy as a shoutrrr service
package ntfy

import (
	"net/http"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service sends notifications Ntfy
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Send a notification message to Ntfy
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendAPI(config *Config, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// apiResponse implements Error

func addHeaderIfNotEmpty(headers *http.Header, key string, value string) {
	_ = "STUB: not implemented"
	return
}
