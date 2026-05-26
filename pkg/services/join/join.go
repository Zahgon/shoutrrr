package join

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	hookURL     = "https://joinjoaomgcd.appspot.com/_ah/api/messaging/v1/sendPush"
	contentType = "text/plain"
)

// Service providing the notification service Pushover
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Send a notification message to Pushover
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendToDevices(devices string, message string, title string, icon string) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}
