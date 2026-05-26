package generators

import (
	"github.com/containrrr/shoutrrr/pkg/generators/basic"
	"github.com/containrrr/shoutrrr/pkg/generators/xouath2"
	"github.com/containrrr/shoutrrr/pkg/services/telegram"
	t "github.com/containrrr/shoutrrr/pkg/types"
)

var generatorMap = map[string]func() t.Generator{
	"basic":    func() t.Generator { return &basic.Generator{} },
	"oauth2":   func() t.Generator { return &xouath2.Generator{} },
	"telegram": func() t.Generator { return &telegram.Generator{} },
}

// NewGenerator creates an instance of the generator that corresponds to the provided identifier
func NewGenerator(identifier string) (t.Generator, error) {
	_ = "STUB: not implemented"
	return *new(t.Generator), nil
}

// ListGenerators lists all available generators
func ListGenerators() []string { _ = "STUB: not implemented"; return nil }
