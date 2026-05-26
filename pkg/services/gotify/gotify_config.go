package gotify

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
)

// Config for use within the gotify plugin
type Config struct {
	standard.EnumlessConfig
	Token      string `url:"path2" desc:"Application token" required:""`
	Host       string `url:"host,port" desc:"Server hostname (and optionally port)" required:""`
	Path       string `optional:"" url:"path1" desc:"Server subpath"`
	Priority   int    `key:"priority" default:"0"`
	Title      string `key:"title" default:"Shoutrrr notification"`
	DisableTLS bool   `key:"disabletls" default:"No"`
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

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "gotify"
)
