package bark

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config for use within the bark service
type Config struct {
	standard.EnumlessConfig
	Title     string `key:"title"    default:""      desc:"Notification title, optionally set by the sender"`
	Host      string `url:"host"                     desc:"Server hostname and port"`
	Path      string `url:"path"     default:"/"     desc:"Server path"`
	DeviceKey string `url:"password"                 desc:"The key for each device"`
	Scheme    string `key:"scheme"   default:"https" desc:"Server protocol, http or https"`
	Sound     string `key:"sound"    default:""      desc:"Value from https://github.com/Finb/Bark/tree/master/Sounds"`
	Badge     int64  `key:"badge"    default:"0"     desc:"The number displayed next to App icon"`
	Icon      string `key:"icon"     default:""      desc:"An url to the icon, available only on iOS 15 or later"`
	Group     string `key:"group"    default:""      desc:"The group of the notification"`
	URL       string `key:"url"      default:""      desc:"Url that will jump when click notification"`
	Category  string `key:"category" default:""      desc:"Reserved field, no use yet"`
	Copy      string `key:"copy"     default:""      desc:"The value to be copied"`
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

// GetAPIURL returns the API URL corresponding to the passed endpoint based on the configuration
func (config *Config) GetAPIURL(endpoint string) string { _ = "STUB: not implemented"; return "" }

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
	Scheme = "bark"
)
