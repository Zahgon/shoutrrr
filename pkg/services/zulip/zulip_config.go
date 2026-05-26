package zulip

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for the zulip service
type Config struct {
	standard.EnumlessConfig
	BotMail string `url:"user" desc:"Bot e-mail address"`
	BotKey  string `url:"pass" desc:"API Key"`
	Host    string `url:"host,port" desc:"API server hostname"`
	Stream  string `key:"stream" optional:"" description:"Target stream name"`
	Topic   string `key:"topic,title" default:""`
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) getURL(_ types.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) setURL(_ types.ConfigQueryResolver, serviceURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// Clone the config to a new Config struct
func (config *Config) Clone() *Config { _ = "STUB: not implemented"; return nil }

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "zulip"
)

// CreateConfigFromURL to use within the zulip service
func CreateConfigFromURL(serviceURL *url.URL) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
