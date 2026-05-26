package join

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for the Pushover notification service service
type Config struct {
	APIKey  string   `url:"pass"`
	Devices []string `key:"devices" desc:"Comma separated list of device IDs"`
	Title   string   `key:"title" optional:"" desc:"If set creates a notification"`
	Icon    string   `key:"icon" optional:"" desc:"Icon URL"`
}

// Enums returns the fields that should use a corresponding EnumFormatter to Print/Parse their values
func (config *Config) Enums() map[string]types.EnumFormatter { _ = "STUB: not implemented"; return nil }

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

// Scheme is the identifying part of this service's configuration URL
const Scheme = "join"
