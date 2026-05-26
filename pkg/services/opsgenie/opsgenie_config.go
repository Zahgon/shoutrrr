package opsgenie

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

const defaultPort = 443

// Config for use within the opsgenie service
type Config struct {
	APIKey      string            `url:"path" desc:"The OpsGenie API key"`
	Host        string            `url:"host" desc:"The OpsGenie API host. Use 'api.eu.opsgenie.com' for EU instances" default:"api.opsgenie.com"`
	Port        uint16            `url:"port" desc:"The OpsGenie API port." default:"443"`
	Alias       string            `key:"alias" desc:"Client-defined identifier of the alert" optional:"true"`
	Description string            `key:"description" desc:"Description field of the alert" optional:"true"`
	Responders  []Entity          `key:"responders" desc:"Teams, users, escalations and schedules that the alert will be routed to send notifications" optional:"true"`
	VisibleTo   []Entity          `key:"visibleTo" desc:"Teams and users that the alert will become visible to without sending any notification" optional:"true"`
	Actions     []string          `key:"actions" desc:"Custom actions that will be available for the alert" optional:"true"`
	Tags        []string          `key:"tags" desc:"Tags of the alert" optional:"true"`
	Details     map[string]string `key:"details" desc:"Map of key-value pairs to use as custom properties of the alert" optional:"true"`
	Entity      string            `key:"entity" desc:"Entity field of the alert that is generally used to specify which domain the Source field of the alert" optional:"true"`
	Source      string            `key:"source" desc:"Source field of the alert" optional:"true"`
	Priority    string            `key:"priority" desc:"Priority level of the alert. Possible values are P1, P2, P3, P4 and P5" optional:"true"`
	Note        string            `key:"note" desc:"Additional note that will be added while creating the alert" optional:"true"`
	User        string            `key:"user" desc:"Display name of the request owner" optional:"true"`
	Title       string            `key:"title" default:"" desc:"notification title, optionally set by the sender"`
}

// Enums returns an empty map because the OpsGenie service doesn't use Enums
func (config Config) Enums() map[string]types.EnumFormatter { _ = "STUB: not implemented"; return nil }

// GetURL is the public version of getURL that creates a new PropKeyResolver when accessed from outside the package
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// Private version of GetURL that can use an instance of PropKeyResolver instead of rebuilding it's model from reflection
func (config *Config) getURL(resolver types.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

// Private version of SetURL that can use an instance of PropKeyResolver instead of rebuilding it's model from reflection
func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "opsgenie"
)
