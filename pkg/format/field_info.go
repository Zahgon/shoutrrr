package format

import (
	r "reflect"

	"github.com/containrrr/shoutrrr/pkg/types"
)

// FieldInfo is the meta data about a config field
type FieldInfo struct {
	Name          string
	Type          r.Type
	EnumFormatter types.EnumFormatter
	Description   string
	DefaultValue  string
	Template      string
	Required      bool
	URLParts      []URLPart
	Title         bool
	Base          int
	Keys          []string
	ItemSeparator rune
}

// IsEnum returns whether a EnumFormatter has been assigned to the field and that it is of a suitable type
func (fi *FieldInfo) IsEnum() bool { _ = "STUB: not implemented"; return false }

// IsURLPart returns whether the field is serialized as the specified part of an URL
func (fi *FieldInfo) IsURLPart(part URLPart) bool { _ = "STUB: not implemented"; return false }

func getStructFieldInfo(structType r.Type, enums map[string]types.EnumFormatter) []FieldInfo {
	_ = "STUB: not implemented"
	return nil
}

// This is an embedded or private field, which should not be part of the Config output

func isHiddenField(field r.StructField) bool { _ = "STUB: not implemented"; return false }

func getFieldBase(field r.StructField) int { _ = "STUB: not implemented"; return 0 }

// Default to base 10 if not tagged
