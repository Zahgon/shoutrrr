package types

// Field is a Key/Value pair used for extra data in log messages
type Field struct {
	Key   string
	Value string
}

// FieldsFromMap creates a Fields slice from a map, optionally sorting keys
func FieldsFromMap(fieldMap map[string]string, sorted bool) []Field {
	_ = "STUB: not implemented"
	return nil
}
