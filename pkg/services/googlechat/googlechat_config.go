package googlechat

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for use within the Google Chat plugin.
type Config struct {
	standard.EnumlessConfig
	Host  string `default:"chat.googleapis.com"`
	Path  string
	Token string
	Key   string
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values.
func (config *Config) setURL(_ types.ConfigQueryResolver, serviceURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func (config *Config) getURL(_ types.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

const (
	// Scheme is the identifying part of this service's configuration URL.
	Scheme = "googlechat"
)
