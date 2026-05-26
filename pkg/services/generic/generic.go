package generic

import (
	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"

	"io"
	"net/url"
)

// Service providing a generic notification service
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Send a notification message to a generic webhook endpoint
func (service *Service) Send(message string, paramsPtr *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a mutable copy of the passed params

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// GetConfigURLFromCustom creates a regular service URL from one with a custom host
func (*Service) GetConfigURLFromCustom(customURL *url.URL) (serviceURL *url.URL, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (service *Service) doSend(config *Config, params types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) getPayload(config *Config, params types.Params) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func createSendParams(config *Config, params types.Params, message string) types.Params {
	_ = "STUB: not implemented"
	return *new(types.Params)
}
