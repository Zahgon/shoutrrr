package format

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// EnumInvalid is the constant value that an enum gets assigned when it could not be parsed
const EnumInvalid = -1

// EnumFormatter is the helper methods for enum-like types
type EnumFormatter struct {
	names       []string
	firstOffset int
	aliases     map[string]int
}

// Names is the list of the valid Enum string values
func (ef EnumFormatter) Names() []string { _ = "STUB: not implemented"; return nil }

// Print takes a enum mapped int and returns it's string representation or "Invalid"
func (ef EnumFormatter) Print(e int) string { _ = "STUB: not implemented"; return "" }

// Parse takes an enum mapped string and returns it's int representation or EnumInvalid (-1)
func (ef EnumFormatter) Parse(s string) int { _ = "STUB: not implemented"; return 0 }

// CreateEnumFormatter creates a EnumFormatter struct
func CreateEnumFormatter(names []string, optAliases ...map[string]int) types.EnumFormatter {
	_ = "STUB: not implemented"
	return *new(types.EnumFormatter)
}
