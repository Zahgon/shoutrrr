package basic

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Generator is the Basic Generator implementation
type Generator struct{}

// Generate generates a service URL from a set of user questions/answers
func (g *Generator) Generate(service types.Service, props map[string]string, _ []string) (types.ServiceConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.ServiceConfig), nil
}

// Clear the property value to skip it next iteration in case of errors

// Handle empty lines

// Use the default value instead of the user input

// No default value is specified, leave it uninitialized
