package slack

import (
	"net/url"
	"regexp"

	"github.com/containrrr/shoutrrr/pkg/types"
)

var _ types.ConfigProp = &Token{}

const (
	hookTokenIdentifier = "hook"
	userTokenIdentifier = "xoxp"
	botTokenIdentifier  = "xoxb"
)

// Token is a Slack API token or a Slack webhook token
type Token struct {
	raw string
}

// SetFromProp updates it's state according to the passed string
// (implementation of the types.ConfigProp interface)
func (token *Token) SetFromProp(propValue string) error { _ = "STUB: not implemented"; return nil }

// GetPropValue returns a deserializable string representation of the token
// (implementation of the types.ConfigProp interface)
func (token *Token) GetPropValue() (string, error) { _ = "STUB: not implemented"; return "", nil }

// TypeIdentifier returns the type identifier of the token
func (token Token) TypeIdentifier() string { _ = "STUB: not implemented"; return "" }

// ParseToken parses and normalizes a token string
func ParseToken(str string) (*Token, error) { _ = "STUB: not implemented"; return nil, nil }

const (
	tokenMatchFull = iota
	tokenMatchType
	tokenMatchPart1
	tokenMatchSep1
	tokenMatchPart2
	tokenMatchSep2
	tokenMatchPart3
	tokenMatchCount
)

var tokenPattern = regexp.MustCompile(`(?:(?P<type>xox.|hook)[-:]|:?)(?P<p1>[A-Z0-9]{9,})(?P<s1>[-/,])(?P<p2>[A-Z0-9]{9,})(?P<s2>[-/,])(?P<p3>[A-Za-z0-9]{24,})`)

// String returns the token in normalized format with dashes (-) as separator
func (token *Token) String() string {
	_ = "STUB: not implemented"

	// UserInfo returns a url.Userinfo struct populated from the token
	return ""
}

func (token *Token) UserInfo() *url.Userinfo { _ = "STUB: not implemented"; return nil }

// IsAPIToken returns whether the identifier is set to anything else but the webhook identifier (`hook`)
func (token *Token) IsAPIToken() bool { _ = "STUB: not implemented"; return false }

const webhookBase = "https://hooks.slack.com/services/"

// WebhookURL returns the corresponding Webhook URL for the Token
func (token Token) WebhookURL() string { _ = "STUB: not implemented"; return "" }

// Authorization returns the corresponding `Authorization` HTTP header value for the Token
func (token *Token) Authorization() string { _ = "STUB: not implemented"; return "" }
