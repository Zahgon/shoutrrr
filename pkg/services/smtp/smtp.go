package smtp

import (
	"io"
	"net/smtp"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/format"
	"github.com/containrrr/shoutrrr/pkg/services/standard"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Service sends notifications to a given e-mail addresses via SMTP
type Service struct {
	standard.Standard
	standard.Templater
	config            *Config
	multipartBoundary string
	propKeyResolver   format.PropKeyResolver
}

const (
	contentHTML      = "text/html; charset=\"UTF-8\""
	contentPlain     = "text/plain; charset=\"UTF-8\""
	contentMultipart = "multipart/alternative; boundary=%s"
)

// Initialize loads ServiceConfig from configURL and sets logger for this Service
func (service *Service) Initialize(configURL *url.URL, logger types.StdLogger) error {
	_ = "STUB: not implemented"
	return nil
}

// Send a notification message to e-mail recipients
func (service *Service) Send(message string, params *types.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func getClientConnection(config *Config) (*smtp.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (service *Service) doSend(client *smtp.Client, message string, config *Config) failure {
	_ = "STUB: not implemented"
	return *new(failure)
}

// Send the QUIT command and close the connection.

func (service *Service) resolveClientHost(config *Config) string {
	_ = "STUB: not implemented"
	return ""
}

func (service *Service) getAuth(config *Config) (smtp.Auth, failure) {
	_ = "STUB: not implemented"
	return *new(smtp.Auth), *new(failure)
}

func (service *Service) sendToRecipient(client *smtp.Client, toAddress string, config *Config, message string) failure {
	_ = "STUB: not implemented"

	// Set the sender and recipient first
	return *new(failure)
}

// Send the email body.

func (service *Service) getHeaders(config *Config, toAddress string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (service *Service) writeMultipartMessage(wc io.WriteCloser, message string) failure {
	_ = "STUB: not implemented"
	return *new(failure)
}

func (service *Service) writeMessagePart(wc io.WriteCloser, message string, template string) failure {
	_ = "STUB: not implemented"
	return *new(failure)
}

func writeMultipartHeader(wc io.WriteCloser, boundary string, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHeaders(wc io.WriteCloser, headers map[string]string) failure {
	_ = "STUB: not implemented"
	return *new(failure)
}
