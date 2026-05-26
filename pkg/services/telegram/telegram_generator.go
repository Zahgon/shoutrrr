package telegram

import (
	"github.com/containrrr/shoutrrr/pkg/types"
	"github.com/containrrr/shoutrrr/pkg/util/generator"

	"io"
)

// Generator is the telegram-specific URL generator
type Generator struct {
	ud        *generator.UserDialog
	client    *Client
	chats     []string
	chatNames []string
	chatTypes []string
	done      bool
	botName   string
	Reader    io.Reader
	Writer    io.Writer
}

// Generate a telegram Shoutrrr configuration from a user dialog
func (g *Generator) Generate(_ types.Service, props map[string]string, _ []string) (types.ServiceConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.ServiceConfig), nil
}

// Subscribe to system signals

// If no updates were retrieved, prompt user to continue

// Another chat was added, prompt user to continue

// Notify API that we got the updates

func (g *Generator) addChat(chat *Chat) (result string) { _ = "STUB: not implemented"; return "" }
