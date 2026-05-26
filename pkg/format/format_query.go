package format

import (
	"net/url"

	t "github.com/containrrr/shoutrrr/pkg/types"
)

// BuildQuery converts the fields of a config object to a delimited query string
func BuildQuery(cqr t.ConfigQueryResolver) string { _ = "STUB: not implemented"; return "" }

// BuildQueryWithCustomFields converts the fields of a config object to a delimited query string,
// escaping any custom fields that share the same key as a config prop using a "__" prefix
func BuildQueryWithCustomFields(cqr t.ConfigQueryResolver, query url.Values) url.Values {
	_ = "STUB: not implemented"
	return *new(url.Values)
}

// Escape any webhook query keys using the same name as service props

// SetConfigPropsFromQuery iterates over all the config prop keys and sets the config prop to the corresponding
// query value based on the key.
// SetConfigPropsFromQuery returns a non-nil url.Values query with all config prop keys removed, even if any of
// them could not be used to set a config field, and with any escaped keys unescaped.
// The error returned is the first error that occurred, subsequent errors are just discarded.
func SetConfigPropsFromQuery(cqr t.ConfigQueryResolver, query url.Values) (url.Values, error) {
	_ = "STUB: not implemented"
	return *new(url.Values), nil
}

// Retrieve the service-related prop value

// Remove it from the query Values

// If an escaped version of the key exist, unescape it

// EscapeKey adds the KeyPrefix to custom URL query keys that conflict with service config prop keys
func EscapeKey(key string) string { _ = "STUB: not implemented"; return "" }

// UnescapeKey removes the KeyPrefix from custom URL query keys that conflict with service config prop keys
func UnescapeKey(key string) string { _ = "STUB: not implemented"; return "" }

// KeyPrefix is the prefix prepended to custom URL query keys that conflict with service config prop keys,
// consisting of two underscore characters ("__")
const KeyPrefix = "__"
