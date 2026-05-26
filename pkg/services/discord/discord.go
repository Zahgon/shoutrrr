package discord

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service providing Discord as a notification service
type Service struct {
	standard.Standard
	config *Config
	pkr    format.PropKeyResolver
}

var limits = types.MessageLimit{
	ChunkSize:      2000,
	TotalChunkSize: 6000,
	ChunkCount:     10,
}

const (
	hookURL = "https://discord.com/api/webhooks"
	// Only search this many runes for a good split position
	maxSearchRunes = 100
)

// Send a notification message to discord
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// SendItems sends items with additional meta data and richer appearance
func (service *Service) SendItems(items []types.MessageItem, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) sendItems(items []types.MessageItem, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateItemsFromPlain creates a set of MessageItems that is compatible with Discords webhook payload
func CreateItemsFromPlain(plain string, splitLines bool) (batches [][]types.MessageItem) {
	_ = "STUB: not implemented"
	return nil
}

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateAPIURLFromConfig takes a discord config object and creates a post url
func CreateAPIURLFromConfig(config *Config) string { _ = "STUB: not implemented"; return "" }

func doSend(payload []byte, postURL string) error { _ = "STUB: not implemented"; return nil }
