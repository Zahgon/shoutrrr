package util

import (
	t "github.com/containrrr/shoutrrr/pkg/types"
)

const ellipsis = " [...]"

// PartitionMessage splits a string into chunks that is at most chunkSize runes, it will search the last distance runes
// for a whitespace to make the split appear nicer. It will keep adding chunks until it reaches maxCount chunks, or if
// the total amount of runes in the chunks reach maxTotal.
// The chunks are returned together with the number of omitted runes (that did not fit into the chunks)
func PartitionMessage(input string, limits t.MessageLimit, distance int) (items []t.MessageItem, omitted int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// If the message is empty, return an empty array

// If no suitable split point is found, use the chunkSize

// ... and start next chunk directly after this one

// The chunk is smaller than the limit, no need to search

// Suitable split point found

// Since the split is on a whitespace, skip it in the next chunk

// Ellipsis returns a string that is at most maxLength characters with a ellipsis appended if the input was longer
func Ellipsis(text string, maxLength int) string { _ = "STUB: not implemented"; return "" }

// MessageItemsFromLines creates a set of MessageItems that is compatible with the supplied limits
func MessageItemsFromLines(plain string, limits t.MessageLimit) (batches [][]t.MessageItem) {
	_ = "STUB: not implemented"
	return nil
}

// Trim and add ellipsis
