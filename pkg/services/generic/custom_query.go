package generic

import (
	"net/url"
)

const extraPrefix = '$'
const headerPrefix = '@'
const caseOffset = 'a' - 'A'

func normalizedHeaderKey(key string) string { _ = "STUB: not implemented"; return "" }

// Char is uppercase

// Add missing dash

// First char, or previous was dash

func appendCustomQueryValues(query url.Values, headers map[string]string, extraData map[string]string) {
	_ = "STUB: not implemented"
	return
}

func stripCustomQueryValues(query url.Values) (headers, extraData map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}
