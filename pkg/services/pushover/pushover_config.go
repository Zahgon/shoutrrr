package pushover

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for the Pushover notification service service
type Config struct {
	Token    string   `url:"pass" desc:"API Token/Key"`
	User     string   `url:"host" desc:"User Key"`
	Devices  []string `key:"devices" optional:""`
	Priority int8     `key:"priority" default:"0"`
	Title    string   `key:"title" optional:""`
}

// Enums returns the fields that should use a corresponding EnumFormatter to Print/Parse their values
func (config *Config) Enums() map[string]types.EnumFormatter { _ = "STUB: not implemented"; return nil }

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// Scheme is the identifying part of this service's configuration URL
const Scheme = "pushover"
