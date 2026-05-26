package shoutrrr

import (
	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/containrrr/shoutrrr/pkg/types"
)

var defaultRouter = router.ServiceRouter{}

// SetLogger sets the logger that the services will use to write progress logs
func SetLogger(logger types.StdLogger) { _ = "STUB: not implemented"; return }

// Send notifications using a supplied url and message
func Send(rawURL string, message string) error { _ = "STUB: not implemented"; return nil }

// CreateSender returns a notification sender configured according to the supplied URL
func CreateSender(rawURLs ...string) (*router.ServiceRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSender returns a notification sender, writing any log output to logger and configured
// to send to the services indicated by the supplied URLs
func NewSender(logger types.StdLogger, serviceURLs ...string) (*router.ServiceRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Version returns the shoutrrr version
func Version() string { _ = "STUB: not implemented"; return "" }
