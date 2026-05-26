package format

import (
	"reflect"

	t "github.com/containrrr/shoutrrr/pkg/types"
)

// PropKeyResolver implements the ConfigQueryResolver interface for services that uses key tags for query props
type PropKeyResolver struct {
	confValue reflect.Value
	keyFields map[string]FieldInfo
	keys      []string
}

// NewPropKeyResolver creates a new PropKeyResolver and initializes it using the provided config
func NewPropKeyResolver(config t.ServiceConfig) PropKeyResolver {
	_ = "STUB: not implemented"
	return *new(PropKeyResolver)
}

// QueryFields returns a list of tagged keys
func (pkr *PropKeyResolver) QueryFields() []string {
	_ = "STUB: not implemented"

	// Get returns the value of a config property tagged with the corresponding key
	return nil
}

func (pkr *PropKeyResolver) Get(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Set sets the value of it's bound struct's property, tagged with the corresponding key
func (pkr *PropKeyResolver) Set(key string, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// set sets the value of a target struct tagged with the corresponding key
func (pkr *PropKeyResolver) set(target reflect.Value, key string, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateConfigFromParams mutates the provided config, updating the values from it's corresponding params
// If the provided config is nil, the internal config will be updated instead.
// The error returned is the first error that occurred, subsequent errors are just discarded.
func (pkr *PropKeyResolver) UpdateConfigFromParams(config t.ServiceConfig, params *t.Params) (firstError error) {
	_ = "STUB: not implemented"
	return nil
}

// SetDefaultProps mutates the provided config, setting the tagged fields with their default values
// If the provided config is nil, the internal config will be updated instead.
// The error returned is the first error that occurred, subsequent errors are just discarded.
func (pkr *PropKeyResolver) SetDefaultProps(config t.ServiceConfig) (firstError error) {
	_ = "STUB: not implemented"
	return nil
}

// Bind is called to create a new instance of the PropKeyResolver, with he internal config reference
// set to the provided config. This should only be used for configs of the same type.
func (pkr *PropKeyResolver) Bind(config t.ServiceConfig) PropKeyResolver {
	_ = "STUB: not implemented"
	return *new(PropKeyResolver)
}

// GetConfigQueryResolver returns the config itself if it implements ConfigQueryResolver
// otherwise it creates and returns a PropKeyResolver that implements it
func GetConfigQueryResolver(config t.ServiceConfig) t.ConfigQueryResolver {
	_ = "STUB: not implemented"
	return *new(t.ConfigQueryResolver)
}

// KeyIsPrimary returns whether the key is the primary (and not an alias)
func (pkr *PropKeyResolver) KeyIsPrimary(key string) bool { _ = "STUB: not implemented"; return false }

func (pkr *PropKeyResolver) configValueOrInternal(config t.ServiceConfig) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func configValue(config t.ServiceConfig) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// IsDefault returns whether the specified key value is the default value
func (pkr *PropKeyResolver) IsDefault(key string, value string) bool {
	_ = "STUB: not implemented"
	return false
}
