package telegram

// Client for Telegram API
type Client struct {
	token string
}

func (c *Client) apiURL(endpoint string) string { _ = "STUB: not implemented"; return "" }

// GetBotInfo returns the bot User info
func (c *Client) GetBotInfo() (*User, error) { _ = "STUB: not implemented"; return nil, nil }

// GetUpdates retrieves the latest updates
func (c *Client) GetUpdates(offset int, limit int, timeout int, allowedUpdates []string) ([]Update, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendMessage sends the specified Message
func (c *Client) SendMessage(message *SendMessagePayload) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetErrorResponse retrieves the error message from a failed request
func GetErrorResponse(body string) error { _ = "STUB: not implemented"; return nil }

// GetResponseError preserves Telegram API errors, falling back to transport errors.
func GetResponseError(err error) error { _ = "STUB: not implemented"; return nil }
