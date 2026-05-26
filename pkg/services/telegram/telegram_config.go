package telegram

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for use within the telegram plugin
type Config struct {
	Token        string    `url:"user"`
	Preview      bool      `key:"preview" default:"Yes" desc:"If disabled, no web page preview will be displayed for URLs"`
	Notification bool      `key:"notification" default:"Yes" desc:"If disabled, sends Message silently"`
	ParseMode    parseMode `key:"parsemode" default:"None" desc:"How the text Message should be parsed"`
	Chats        []string  `key:"chats,channels" desc:"Chat IDs or Channel names (using @channel-name)"`
	Title        string    `key:"title" default:"" desc:"Notification title, optionally set by the sender"`
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
const (
	Scheme = "telegram"
)
