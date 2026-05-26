package matrix

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	t "github.com/containrrr/shoutrrr/pkg/types"
)

const defaultDeviceID = "shoutrrr"

// Config is the configuration for the matrix service
type Config struct {
	standard.EnumlessConfig

	User       string   `optional:"" url:"user" desc:"Username or empty when using access token"`
	Password   string   `url:"password" desc:"Password or access token"`
	DisableTLS bool     `key:"disableTLS" default:"No"`
	DeviceID   string   `key:"deviceID" default:"shoutrrr" desc:"Device ID for password login; keeps Matrix homeservers from creating a new device for each login"`
	Host       string   `url:"host"`
	Rooms      []string `key:"rooms,room" optional:"" desc:"Room aliases, or with ! prefix, room IDs"`
	Title      string   `key:"title" default:""`
}

// GetURL returns a URL representation of it's current field values
func (c *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (c *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (c *Config) getURL(resolver t.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) setURL(resolver t.ConfigQueryResolver, configURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// If room does not begin with a '#' let's prepend it
