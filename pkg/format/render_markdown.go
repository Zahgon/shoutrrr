package format

import (
	"strings"
)

// MarkdownTreeRenderer renders a ContainerNode tree into a markdown documentation string
type MarkdownTreeRenderer struct {
	HeaderPrefix      string
	PropsDescription  string
	PropsEmptyMessage string
}

// RenderTree renders a ContainerNode tree into a markdown documentation string
func (r MarkdownTreeRenderer) RenderTree(root *ContainerNode, scheme string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r MarkdownTreeRenderer) writeURLFields(sb *strings.Builder, urlFields []*FieldInfo, scheme string) {
	_ = "STUB: not implemented"
	return
}

// Host cannot be empty

// Hard coded override for host:port 😓

func (MarkdownTreeRenderer) writeFieldExtras(sb *strings.Builder, field *FieldInfo) {
	_ = "STUB: not implemented"
	return
}

// Skip primary alias (as it's the same as the field name)

func (MarkdownTreeRenderer) writeFieldPrimary(sb *strings.Builder, field *FieldInfo) {
	_ = "STUB: not implemented"
	return
}

func (r MarkdownTreeRenderer) writeHeader(sb *strings.Builder, text string) {
	_ = "STUB: not implemented"
	return
}
