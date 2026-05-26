package util

import (
	"io/ioutil"
	"log"
)

// Min returns the smallest of a and b
func Min(a int, b int) int { _ = "STUB: not implemented"; return 0 }

// Max returns the largest of a and b
func Max(a int, b int) int { _ = "STUB: not implemented"; return 0 }

// DiscardLogger is a logger that discards any output written to it
var DiscardLogger = log.New(ioutil.Discard, "", 0)
