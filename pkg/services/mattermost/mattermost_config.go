package mattermost

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config object holding all information
type Config struct {
	standard.EnumlessConfig
	UserName string `url:"user" optional:"" desc:"Override webhook user"`
	Icon     string `key:"icon,icon_emoji,icon_url" default:"" optional:"" desc:"Use emoji or URL as icon (based on presence of http(s):// prefix)"`
	Title    string `key:"title" default:"" desc:"Notification title, optionally set by the sender (not used)"`
	Channel  string `url:"path2" optional:"" desc:"Override webhook channel"`
	Host     string `url:"host,port" desc:"Mattermost server host"`
	Token    string `url:"path1" desc:"Webhook token"`
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) setURL(resolver types.ConfigQueryResolver, serviceURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorMessage for error events within the mattermost service
type ErrorMessage string

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "mattermost"
	// NotEnoughArguments provided in the service URL
	NotEnoughArguments ErrorMessage = "the apiURL does not include enough arguments, either provide 1 or 3 arguments (they may be empty)"
)

// CreateConfigFromURL to use within the mattermost service
func CreateConfigFromURL(serviceURL *url.URL) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
