package ifttt

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "ifttt"
)

// Config is the configuration needed to send IFTTT notifications
type Config struct {
	standard.EnumlessConfig
	WebHookID         string   `url:"host" required:"true"`
	Events            []string `key:"events" required:"true"`
	Value1            string   `key:"value1" optional:""`
	Value2            string   `key:"value2" optional:""`
	Value3            string   `key:"value3" optional:""`
	UseMessageAsValue uint8    `key:"messagevalue" desc:"Sets the corresponding value field to the notification message" default:"2"`
	UseTitleAsValue   uint8    `key:"titlevalue" desc:"Sets the corresponding value field to the notification title" default:"0"`
	Title             string   `key:"title" default:"" desc:"Notification title, optionally set by the sender"`
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
