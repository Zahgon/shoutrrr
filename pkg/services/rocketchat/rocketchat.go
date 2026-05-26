package rocketchat

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service sends notifications to a pre-configured channel or user
type Service struct {
	standard.Standard
	config *Config
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a notification message to Rocket.chat
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func buildURL(config *Config) string { _ = "STUB: not implemented"; return "" }
