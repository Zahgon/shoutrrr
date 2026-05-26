package rocketchat

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// JSON used within the Rocket.chat service
type JSON struct {
	Text     string `json:"text"`
	UserName string `json:"username,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// CreateJSONPayload compatible with the rocket.chat webhook api
func CreateJSONPayload(config *Config, message string, params *types.Params) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
