package pushbullet

import (
	"errors"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"

	"github.com/containrrr/shoutrrr/pkg/services/standard"
)

// Config ...
type Config struct {
	standard.EnumlessConfig
	Targets []string `url:"path"`
	Token   string   `url:"host"`
	Title   string   `key:"title" default:"Shoutrrr notification"`
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

// Remove initial slash to skip empty first target

func validateToken(token string) error { _ = "STUB: not implemented"; return nil }

const (
	//Scheme is the scheme part of the service configuration URL
	Scheme = "pushbullet"
)

// ErrorTokenIncorrectSize is the error returned when the token size is incorrect
var ErrorTokenIncorrectSize = errors.New("token has incorrect size")
