package mattermost

import (
	"regexp"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// JSON payload for mattermost notifications
type JSON struct {
	Text      string `json:"text"`
	UserName  string `json:"username,omitempty"`
	Channel   string `json:"channel,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
}

var iconURLPattern = regexp.MustCompile(`https?://`)

// SetIcon sets the appropriate icon field in the payload based on whether the input is a URL or not
func (j *JSON) SetIcon(icon string) { _ = "STUB: not implemented"; return }

// CreateJSONPayload for usage with the mattermost service
func CreateJSONPayload(config *Config, message string, params *types.Params) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
