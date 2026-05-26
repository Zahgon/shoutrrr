package format

// ParseBool returns true for "1","true","yes" or false for "0","false","no" or defaultValue for any other value
func ParseBool(value string, defaultValue bool) (parsedValue bool, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

// PrintBool returns "Yes" if value is true, otherwise returns "No"
func PrintBool(value bool) string { _ = "STUB: not implemented"; return "" }

// IsNumber returns whether the specified string is number-like
func IsNumber(value string) bool { _ = "STUB: not implemented"; return false }
