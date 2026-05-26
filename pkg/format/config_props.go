package format

import (
	r "reflect"
)

// GetConfigPropFromString deserializes a config property from a string representation using the ConfigProp interface
func GetConfigPropFromString(structType r.Type, value string) (r.Value, error) {
	_ = "STUB: not implemented"
	return *new(r.Value), nil
}

// GetConfigPropString serializes a config property to a string representation using the ConfigProp interface
func GetConfigPropString(propPtr r.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
