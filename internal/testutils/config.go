package testutils

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// TestConfigGetInvalidQueryValue tests whether the config returns an error when an invalid query value is requested
func TestConfigGetInvalidQueryValue(config types.ServiceConfig) { _ = "STUB: not implemented"; return }

// TestConfigSetInvalidQueryValue tests whether the config returns an error when a URL with an invalid query value is parsed
func TestConfigSetInvalidQueryValue(config types.ServiceConfig, rawInvalidURL string) {
	_ = "STUB: not implemented"
	return
}

// TestConfigSetDefaultValues tests whether setting the default values can be set for an empty config without any errors
func TestConfigSetDefaultValues(config types.ServiceConfig) { _ = "STUB: not implemented"; return }

// TestConfigGetEnumsCount tests whether the config.Enums returns the expected amount of items
func TestConfigGetEnumsCount(config types.ServiceConfig, expectedCount int) {
	_ = "STUB: not implemented"
	return
}

// TestConfigGetFieldsCount tests whether the config.QueryFields return the expected amount of fields
func TestConfigGetFieldsCount(config types.ServiceConfig, expectedCount int) {
	_ = "STUB: not implemented"
	return
}
