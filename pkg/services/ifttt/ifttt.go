package ifttt

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	apiURLFormat = "https://maker.ifttt.com/trigger/%s/with/key/%s"
)

// Service sends notifications to a IFTTT webhook
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a notification message to a IFTTT webhook
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateAPIURLForEvent creates a IFTTT webhook URL for the given event
func (service *Service) createAPIURLForEvent(event string) string {
	_ = "STUB: not implemented"
	return ""
}

func doSend(payload []byte, postURL string) error { _ = "STUB: not implemented"; return nil }
