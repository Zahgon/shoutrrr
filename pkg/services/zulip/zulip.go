package zulip

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

const (
	contentMaxSize = 10000 // bytes
	topicMaxLength = 60    // characters
)

// Send a notification message to Zulip
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	// Clone the config because we might modify stream and/or
	// topic with values from the parameters and they should only
	// change this Send().
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) doSend(config *Config, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) getAPIURL(config *Config) string { _ = "STUB: not implemented"; return "" }
