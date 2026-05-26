package zulip

import (
	"net/url"
)

// CreatePayload compatible with the zulip api
func CreatePayload(config *Config, message string) url.Values {
	_ = "STUB: not implemented"
	return *new(url.Values)
}
