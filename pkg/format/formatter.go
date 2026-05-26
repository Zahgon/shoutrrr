package format

import (
	r "reflect"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// GetServiceConfig returns the inner config of a service
func GetServiceConfig(service types.Service) types.ServiceConfig {
	_ = "STUB: not implemented"
	return *new(types.ServiceConfig)
}

//

// ColorFormatTree returns a color highlighted string representation of a node tree
func ColorFormatTree(rootNode *ContainerNode, withValues bool) string {
	_ = "STUB: not implemented"
	return ""
}

// GetServiceConfigFormat returns type and field information about a ServiceConfig, resolved from it's Service
func GetServiceConfigFormat(service types.Service) *ContainerNode {
	_ = "STUB: not implemented"
	return nil
}

// GetConfigFormat returns type and field information about a ServiceConfig
func GetConfigFormat(config types.ServiceConfig) *ContainerNode {
	_ = "STUB: not implemented"
	return nil

	// SetConfigField deserializes the inputValue and sets the field of a config to that value
}

func SetConfigField(config r.Value, field FieldInfo, inputValue string) (valid bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// When updating a pointer slice, use the value type kind

// If not a pointer slice, dereference the value

// Use the split string parts as the target value

func getMapValue(valueType r.Type, valueRaw string) (r.Value, error) {
	_ = "STUB: not implemented"
	return *new(r.Value), nil
}

func getMapUintValue(valueRaw string, bits int, kind r.Kind) (r.Value, error) {
	_ = "STUB: not implemented"
	return *new(r.Value), nil
}

func getMapIntValue(valueRaw string, bits int, kind r.Kind) (r.Value, error) {
	_ = "STUB: not implemented"
	return *new(r.Value), nil
}

// GetConfigFieldString serializes the config field value to a string representation
func GetConfigFieldString(config r.Value, field FieldInfo) (value string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
