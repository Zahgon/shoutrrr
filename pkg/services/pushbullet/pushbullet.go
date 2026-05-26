package pushbullet

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/containrrr/shoutrrr/pkg/util/jsonclient"
)

const (
	pushesEndpoint = "https://api.pushbullet.com/v2/pushes"
)

// Service providing Pushbullet as a notification service
type Service struct {
	standard.Standard
	client jsonclient.Client
	config *Config
	pkr    format.PropKeyResolver
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a push notification via Pushbullet
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func doSend(config *Config, target string, message string, client jsonclient.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Look at response fields?
