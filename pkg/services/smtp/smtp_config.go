package smtp

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// Config is the configuration needed to send e-mail notifications over SMTP
type Config struct {
	Host        string    `desc:"SMTP server hostname or IP address" url:"Host"`
	Username    string    `desc:"SMTP server username" default:"" url:"User"`
	Password    string    `desc:"SMTP server password or hash (for OAuth2)" default:"" url:"Pass"`
	Port        uint16    `desc:"SMTP server port, common ones are 25, 465, 587 or 2525" default:"25" url:"Port"`
	FromAddress string    `desc:"E-mail address that the mail are sent from" key:"fromaddress,from"`
	FromName    string    `desc:"Name of the sender" optional:"yes" key:"fromname"`
	ToAddresses []string  `desc:"List of recipient e-mails separated by \",\" (comma)" key:"toaddresses,to"`
	Subject     string    `desc:"The subject of the sent mail" key:"subject,title" default:"Shoutrrr Notification"`
	Auth        authType  `desc:"SMTP authentication method" key:"auth" default:"Unknown"`
	Encryption  encMethod `desc:"Encryption method" default:"Auto" key:"encryption"`
	UseStartTLS bool      `desc:"Whether to use StartTLS encryption" default:"Yes" key:"usestarttls,starttls"`
	UseHTML     bool      `desc:"Whether the message being sent is in HTML" default:"No" key:"usehtml"`
	ClientHost  string    `desc:"The client host name sent to the SMTP server during HELLO phase. If set to \"auto\" it will use the OS hostname" key:"clienthost" default:"localhost"`
}

// GetURL returns a URL representation of its current field values
func (config *Config) GetURL() *url.URL { _ = "STUB: not implemented"; return nil }

// SetURL updates a ServiceConfig from a URL representation of its field values
func (config *Config) SetURL(url *url.URL) error { _ = "STUB: not implemented"; return nil }

func (config *Config) getURL(resolver types.ConfigQueryResolver) *url.URL {
	_ = "STUB: not implemented"
	return nil
}

func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

// Clone returns a copy of the config
func (config *Config) Clone() Config { _ = "STUB: not implemented"; return *new(Config) }

// FixEmailTags replaces parsed spaces (+) in e-mail addresses with '+'
func (config *Config) FixEmailTags() { _ = "STUB: not implemented"; return }

// Enums returns the fields that should use a corresponding EnumFormatter to Print/Parse their values
func (config *Config) Enums() map[string]types.EnumFormatter { _ = "STUB: not implemented"; return nil }

// Scheme is the identifying part of this service's configuration URL
const Scheme = "smtp"
