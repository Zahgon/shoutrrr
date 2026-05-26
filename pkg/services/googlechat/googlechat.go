package googlechat

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service providing Google Chat as a notification service.
type Service struct {
	standard.Standard
	config *Config
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service.
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a notification message to Google Chat.
func (service *Service) Send(message string, _ *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func getAPIURL(config *Config) *url.URL { _ = "STUB: not implemented"; return nil }
