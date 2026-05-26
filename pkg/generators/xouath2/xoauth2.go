package xouath2

import (
	"github.com/containrrr/shoutrrr/pkg/services/smtp"
	"github.com/containrrr/shoutrrr/pkg/types"
	"golang.org/x/oauth2"
)

// Generator is the XOAuth2 Generator implementation
type Generator struct{}

// Generate generates a service URL from a set of user questions/answers
func (g *Generator) Generate(_ types.Service, props map[string]string, args []string) (types.ServiceConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.ServiceConfig), nil
}

func oauth2GeneratorFile(file string) (*smtp.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func oauth2Generator() (*smtp.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func oauth2GeneratorGmail(credFile string) (*smtp.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateOauth2Config(conf *oauth2.Config, host string) (*smtp.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
