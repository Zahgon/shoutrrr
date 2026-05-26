package telegram

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	apiFormat = "https://api.telegram.org/bot%s/%s"
	maxlength = 4096
)

// Service sends notifications to a given telegram chat
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

// Send notification to Telegram
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendMessageForChatIDs(message string, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// GetConfig returns the Config for the service
func (service *Service) GetConfig() *Config { _ = "STUB: not implemented"; return nil }

func sendMessageToAPI(message string, chat string, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}
