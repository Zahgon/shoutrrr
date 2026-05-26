package discord

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config is the configuration needed to send discord notifications
type Config struct {
	standard.EnumlessConfig
	WebhookID  string `url:"host"`
	Token      string `url:"user"`
	Title      string `key:"title"      default:""`
	Username   string `key:"username"   default:""         desc:"Override the webhook default username"`
	Avatar     string `key:"avatar,avatarurl"     default:""         desc:"Override the webhook default avatar with specified URL"`
	Color      uint   `key:"color"      default:"0x50D9ff" desc:"The color of the left border for plain messages"   base:"16"`
	ColorError uint   `key:"colorError" default:"0xd60510" desc:"The color of the left border for error messages"   base:"16"`
	ColorWarn  uint   `key:"colorWarn"  default:"0xffc441" desc:"The color of the left border for warning messages" base:"16"`
	ColorInfo  uint   `key:"colorInfo"  default:"0x2488ff" desc:"The color of the left border for info messages"    base:"16"`
	ColorDebug uint   `key:"colorDebug" default:"0x7b00ab" desc:"The color of the left border for debug messages"   base:"16"`
	SplitLines bool   `key:"splitLines" default:"Yes"      desc:"Whether to send each line as a separate embedded item"`
	JSON       bool   `key:"json"       default:"No"       desc:"Whether to send the whole message as the JSON payload instead of using it as the 'content' field"`
	ThreadID   string `key:"thread_id"  default:""         desc:"Optional thread ID for posting into a channel thread"`
}

// LevelColors returns an array of colors with a MessageLevel index
func (config *Config) LevelColors() (colors [types.MessageLevelCount]uint) {
	_ = "STUB: not implemented"
	return nil
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) getURL(resolver types.ConfigQueryResolver) (u *url.URL) {
	_ = "STUB: not implemented"
	return nil
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// Scheme is the identifying part of this service's configuration URL
const Scheme = "discord"
