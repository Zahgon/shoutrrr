package format

import (
	r "reflect"
)

// NodeTokenType is used to represent the type of value that a node has for syntax highlighting
type NodeTokenType int

const (
	// UnknownToken represents all unknown/unspecified tokens
	UnknownToken NodeTokenType = iota
	// NumberToken represents all numbers
	NumberToken
	// StringToken represents strings and keys
	StringToken
	// EnumToken represents enum values
	EnumToken
	// TrueToken represent boolean true
	TrueToken
	// FalseToken represent boolean false
	FalseToken
	// PropToken represent a serializable struct prop
	PropToken
	// ErrorToken represent a value that was not serializable or otherwise invalid
	ErrorToken
	// ContainerToken is used for Array/Slice and Map tokens
	ContainerToken
)

// Node is the generic config tree item
type Node interface {
	Field() *FieldInfo
	TokenType() NodeTokenType
	Update(tv r.Value)
}

// ValueNode is a Node without any child items
type ValueNode struct {
	*FieldInfo
	Value     string
	tokenType NodeTokenType
}

// Field returns the inner FieldInfo
func (n *ValueNode) Field() *FieldInfo {
	_ = "STUB: not implemented"

	// TokenType returns a NodeTokenType that matches the value
	return nil
}

func (n *ValueNode) TokenType() NodeTokenType {
	_ = "STUB: not implemented"
	return *

	// Update updates the value string from the provided value
	new(NodeTokenType)
}

func (n *ValueNode) Update(tv r.Value) { _ = "STUB: not implemented"; return }

// ContainerNode is a Node with child items
type ContainerNode struct {
	*FieldInfo
	Items        []Node
	MaxKeyLength int
}

// Field returns the inner FieldInfo
func (n *ContainerNode) Field() *FieldInfo {
	_ = "STUB: not implemented"

	// TokenType always returns ContainerToken for ContainerNode
	return nil
}

func (n *ContainerNode) TokenType() NodeTokenType {
	_ = "STUB: not implemented"
	return *

	// Update updates the items to match the provided value
	new(NodeTokenType)
}

func (n *ContainerNode) Update(tv r.Value) { _ = "STUB: not implemented"; return }

func (n *ContainerNode) updateArrayNode(arrayValue r.Value) { _ = "STUB: not implemented"; return }

func getArrayNode(arrayValue r.Value, fieldInfo *FieldInfo) (node *ContainerNode) {
	_ = "STUB: not implemented"
	return nil
}

func sortNodeItems(nodeItems []Node) { _ = "STUB: not implemented"; return }

func (n *ContainerNode) updateMapNode(mapValue r.Value) { _ = "STUB: not implemented"; return }

// The keys will always be strings

func getMapNode(mapValue r.Value, fieldInfo *FieldInfo) (node *ContainerNode) {
	_ = "STUB: not implemented"
	return nil
}

func getNode(fieldVal r.Value, fieldInfo *FieldInfo) Node {
	_ = "STUB: not implemented"
	return *new(Node)
}

func getRootNode(v interface{}) *ContainerNode { _ = "STUB: not implemented"; return nil }

// The current field is Anonymous and not present in the FieldInfo slice

func getValueNode(fieldVal r.Value, fieldInfo *FieldInfo) (node *ValueNode) {
	_ = "STUB: not implemented"
	return nil
}

func getValueNodeValue(fieldValue r.Value, fieldInfo *FieldInfo) (string, NodeTokenType) {
	_ = "STUB: not implemented"
	return "", *new(NodeTokenType)
}

// Unsupported value

func getContainerValueString(fieldValue r.Value, fieldInfo *FieldInfo) string {
	_ = "STUB: not implemented"
	return ""
}

// Inherit the base from the container
