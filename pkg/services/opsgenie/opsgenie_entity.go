package opsgenie

// Entity represents either a user or a team
//
// The different variations are:
//
// { "id":"4513b7ea-3b91-438f-b7e4-e3e54af9147c", "type":"team" }
// { "name":"rocket_team", "type":"team" }
// { "id":"bb4d9938-c3c2-455d-aaab-727aa701c0d8", "type":"user" }
// { "username":"trinity@opsgenie.com", "type":"user" }
type Entity struct {
	Type     string `json:"type"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
}

// SetFromProp deserializes an entity
func (e *Entity) SetFromProp(propValue string) error { _ = "STUB: not implemented"; return nil }

// GetPropValue serializes an entity
func (e *Entity) GetPropValue() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Detects OpsGenie IDs in the form 4513b7ea-3b91-438f-b7e4-e3e54af9147c
func isOpsGenieID(str string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
