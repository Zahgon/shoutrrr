package teams

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
)

// Config for use within the teams plugin
type Config struct {
	standard.EnumlessConfig
	Group      string `url:"user" optional:""`
	Tenant     string `url:"host" optional:""`
	AltID      string `url:"path1" optional:""`
	GroupOwner string `url:"path2" optional:""`

	Title string `key:"title" optional:""`
	Color string `key:"color" optional:""`
	Host  string `key:"host" optional:"" default:"outlook.office.com"`
}

func (config *Config) webhookParts() [4]string { _ = "STUB: not implemented"; return nil }

// SetFromWebhookURL updates the config WebhookParts from a teams webhook URL
func (config *Config) SetFromWebhookURL(webhookURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigFromWebhookURL creates a new Config from a parsed Teams Webhook URL
func ConfigFromWebhookURL(webhookURL url.URL) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) getURL(resolver types.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (config *Config) setFromWebhookParts(parts [4]string) { _ = "STUB: not implemented"; return }

func buildWebhookURL(host, group, tenant, altID, groupOwner string) string {
	_ = "STUB: not implemented"
	// config.Group, config.Tenant, config.AltID, config.GroupOwner
	return ""
}

func parseAndVerifyWebhookURL(webhookURL string) (parts [4]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "teams"
	// LegacyHost is the default host for legacy webhook requests
	LegacyHost = "outlook.office.com"
	// LegacyPath is the initial path of the webhook URL for legacy webhook requests
	LegacyPath = "webhook"
	// Path is the initial path of the webhook URL for domain-scoped webhook requests
	Path = "webhookb2"
	// ProviderName is the name of the Teams integration provider
	ProviderName = "IncomingWebhook"
)
