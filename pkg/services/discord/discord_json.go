package discord

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// WebhookPayload is the webhook endpoint payload
type WebhookPayload struct {
	Embeds    []embedItem `json:"embeds"`
	Username  string      `json:"username,omitempty"`
	AvatarURL string      `json:"avatar_url,omitempty"`
}

// JSON is the actual notification payload
type embedItem struct {
	Title     string       `json:"title,omitempty"`
	Content   string       `json:"description,omitempty"`
	URL       string       `json:"url,omitempty"`
	Timestamp string       `json:"timestamp,omitempty"`
	Color     uint         `json:"color,omitempty"`
	Footer    *embedFooter `json:"footer,omitempty"`
}

type embedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

// CreatePayloadFromItems creates a JSON payload to be sent to the discord webhook API
func CreatePayloadFromItems(items []types.MessageItem, title string, colors [types.MessageLevelCount]uint) (WebhookPayload, error) {
	_ = "STUB: not implemented"
	return *new(WebhookPayload), nil
}

// This should not happen, but it's better to leave the index check before dereferencing the array
