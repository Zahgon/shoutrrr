package router

import (
	"net/url"
	"time"

	t "github.com/containrrr/shoutrrr/pkg/types"
)

// ServiceRouter is responsible for routing a message to a specific notification service using the notification URL
type ServiceRouter struct {
	logger   t.StdLogger
	services []t.Service
	queue    []string
	Timeout  time.Duration
}

// New creates a new service router using the specified logger and service URLs
func New(logger t.StdLogger, serviceURLs ...string) (*ServiceRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddService initializes the specified service from its URL, and adds it if no errors occur
func (router *ServiceRouter) AddService(serviceURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// Send sends the specified message using the routers underlying services
func (router *ServiceRouter) Send(message string, params *t.Params) []error {
	_ = "STUB: not implemented"
	return nil
}

// SendItems sends the specified message items using the routers underlying services
func (router *ServiceRouter) SendItems(items []t.MessageItem, params t.Params) []error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback using old API for now

// SendAsync sends the specified message using the routers underlying services
func (router *ServiceRouter) SendAsync(message string, params *t.Params) chan error {
	_ = "STUB: not implemented"
	return nil
}

func sendToService(service t.Service, results chan error, timeout time.Duration, message string, params t.Params) {
	_ = "STUB: not implemented"
	return
}

// TODO: There really ought to be a better way to name the services

// Enqueue adds the message to an internal queue and sends it when Flush is invoked
func (router *ServiceRouter) Enqueue(message string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Flush sends all messages that have been queued up as a combined message. This method should be deferred!
func (router *ServiceRouter) Flush(params *t.Params) {
	_ = "STUB: not implemented"
	// Since this method is supposed to be deferred we just have to ignore errors
	return
}

// SetLogger sets the logger that the services will use to write progress logs
func (router *ServiceRouter) SetLogger(logger t.StdLogger) { _ = "STUB: not implemented"; return }

// ExtractServiceName from a notification URL
func (router *ServiceRouter) ExtractServiceName(rawURL string) (string, *url.URL, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Route a message to a specific notification service using the notification URL
func (router *ServiceRouter) Route(rawURL string, message string) error {
	_ = "STUB: not implemented"
	return nil
}

func (router *ServiceRouter) initService(rawURL string) (t.Service, error) {
	_ = "STUB: not implemented"
	return *new(t.Service), nil
}

// NewService returns a new uninitialized service instance
func (*ServiceRouter) NewService(serviceScheme string) (t.Service, error) {
	_ = "STUB: not implemented"
	return *new(t.Service), nil
}

// newService returns a new uninitialized service instance
func newService(serviceScheme string) (t.Service, error) {
	_ = "STUB: not implemented"
	return *new(t.Service), nil
}

// ListServices returns the available services
func (router *ServiceRouter) ListServices() []string { _ = "STUB: not implemented"; return nil }

// Locate returns the service implementation that corresponds to the given service URL
func (router *ServiceRouter) Locate(rawURL string) (t.Service, error) {
	_ = "STUB: not implemented"
	return *new(t.Service), nil
}

func (router *ServiceRouter) log(v ...interface{}) { _ = "STUB: not implemented"; return }
