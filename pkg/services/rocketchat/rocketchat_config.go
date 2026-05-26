package rocketchat

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
)

// Config for the rocket.chat service
type Config struct {
	standard.EnumlessConfig
	UserName string `url:"user" optional:""`
	Host     string `url:"host"`
	Port     string `url:"port"`
	TokenA   string `url:"path1"`
	Channel  string `url:"path3"`
	TokenB   string `url:"path2"`
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(serviceURL *url.URL) error { _ = "STUB: not implemented"; return nil }

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "rocketchat"
	// NotEnoughArguments provided in the service URL
	NotEnoughArguments = "the apiURL does not include enough arguments"
)

// CreateConfigFromURL to use within the rocket.chat service
func CreateConfigFromURL(_ types.ConfigQueryResolver, serviceURL *url.URL) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
