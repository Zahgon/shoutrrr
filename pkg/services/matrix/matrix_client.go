package matrix

import (
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/types"
)

type client struct {
	apiURL      url.URL
	accessToken string
	logger      types.StdLogger
	txnID       uint64
}

func newClient(host string, disableTLS bool, logger types.StdLogger) (c *client) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) useToken(token string) { _ = "STUB: not implemented"; return }

func (c *client) login(user string, password string, deviceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) loginPassword(user string, password string, deviceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) sendMessage(message string, rooms []string) (errors []error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) sendToExplicitRooms(rooms []string, message string) (errors []error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) sendToJoinedRooms(message string) (errors []error) {
	_ = "STUB: not implemented"
	return nil
}

// Send to all rooms that are joined

func (c *client) joinRoom(room string) (roomID string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *client) sendMessageToRoom(message string, roomID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) apiGet(path string, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) apiPost(path string, request interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) apiPut(path string, request interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) apiRequest(method string, path string, request interface{}, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) nextTransactionID() string { _ = "STUB: not implemented"; return "" }

func (c *client) updateAccessToken() { _ = "STUB: not implemented"; return }

func (c *client) logf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (c *client) getJoinedRooms() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
