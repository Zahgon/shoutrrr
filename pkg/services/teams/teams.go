package teams

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service providing teams as a notification service
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Send a notification message to Microsoft Teams
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

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

func (service *Service) doSend(config *Config, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// Teams need a summary for the webhook, use title or first (truncated) row

// Emit a warning to the log for now.
// TODO(v0.6): Remove legacy support as it should be fully deprecated now
