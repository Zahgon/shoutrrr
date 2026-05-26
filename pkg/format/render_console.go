package format

import (
	"strings"
)

// ConsoleTreeRenderer renders a ContainerNode tree into a ansi-colored console string
type ConsoleTreeRenderer struct {
	WithValues bool
}

// RenderTree renders a ContainerNode tree into a ansi-colored console string
func (r ConsoleTreeRenderer) RenderTree(root *ContainerNode, _ string) string {
	_ = "STUB: not implemented"
	return ""
}

// Since no values was supplied, let's substitute the value with the type

// If the value is an enum type, providing the name is a bit pointless
// Instead, use a common string "option" to signify the type

// Skip primary alias (as it's the same as the field name)

func (r ConsoleTreeRenderer) writeNodeValue(sb *strings.Builder, node Node) int {
	_ = "STUB: not implemented"
	return 0
}

func (r ConsoleTreeRenderer) writeContainer(sb *strings.Builder, node *ContainerNode) int {
	_ = "STUB: not implemented"
	return 0
}
