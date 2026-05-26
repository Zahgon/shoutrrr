package matrix

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	t "github.com/containrrr/shoutrrr/pkg/types"
)

// Scheme is the identifying part of this service's configuration URL
const Scheme = "matrix"

// Service providing Matrix as a notification service
type Service struct {
	standard.Standard
	config *Config
	client *client
	pkr    format.PropKeyResolver
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (s *Service) Initialize(configURL *url.URL, logger t.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send notification
func (s *Service) Send(message string, params *t.Params) error {
	_ = "STUB: not implemented"
	return nil
}
