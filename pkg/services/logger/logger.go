package logger

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service is the Logger service struct
type Service struct {
	standard.Standard
	config *Config
}

// Send a notification message to log
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) doSend(data types.Params) error { _ = "STUB: not implemented"; return nil }

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(_ *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}
