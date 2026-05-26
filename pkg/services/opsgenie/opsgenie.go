package opsgenie

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	alertEndpointTemplate = "https://%s:%d/v2/alerts"
)

// Service providing OpsGenie as a notification service
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

func (service *Service) sendAlert(url string, apiKey string, payload AlertPayload) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a notification message to OpsGenie
// See: https://docs.opsgenie.com/docs/alert-api#create-alert
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) newAlertPayload(message string, params *types.Params) (AlertPayload, error) {
	_ = "STUB: not implemented"
	return *new(AlertPayload), nil
}

// Defensive copy

// Use `Message` for the title if available, or if the message is too long
// Use `Description` for the message in these scenarios
