package slack

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service sends notifications to a pre-configured channel or user
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

const (
	apiPostMessage = "https://slack.com/api/chat.postMessage"
)

// Send a notification message to Slack
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendAPI(config *Config, payload interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendWebhook(config *Config, payload interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Treat status 200 as no error regardless of actual content
