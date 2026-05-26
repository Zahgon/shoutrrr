package ifttt

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// jsonPayload is the actual notification payload
type jsonPayload struct {
	Value1 string `json:"value1" `
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

// createJSONToSend creates a jsonPayload payload to be sent to the IFTTT webhook API
func createJSONToSend(config *Config, message string, params *types.Params) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
