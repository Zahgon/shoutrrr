package generator

import (
	"bufio"
	"errors"
	"io"
	re "regexp"
)

var errInvalidFormat = errors.New("invalid format")

// ValidateFormat is a validation wrapper turning false bool results into errors
func ValidateFormat(validator func(string) bool) func(string) error {
	_ = "STUB: not implemented"
	return nil
}

var errRequired = errors.New("field is required")

// Required is a validator that checks whether the input contains any characters
func Required(answer string) error { _ = "STUB: not implemented"; return nil }

// UserDialog is an abstraction for question/answer based user interaction
type UserDialog struct {
	reader  io.Reader
	writer  io.Writer
	scanner *bufio.Scanner
	props   map[string]string
}

// NewUserDialog initializes a UserDialog with safe defaults
func NewUserDialog(reader io.Reader, writer io.Writer, props map[string]string) *UserDialog {
	_ = "STUB: not implemented"
	return nil
}

// Write message to user
func (ud *UserDialog) Write(message string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Writeln writes a message to the user that completes a line
func (ud *UserDialog) Writeln(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Query writes the prompt to the user and returns the regex groups if it matches the validator pattern
func (ud *UserDialog) Query(prompt string, validator *re.Regexp, key string) (groups []string) {
	_ = "STUB: not implemented"
	return nil
}

// QueryAll is a version of Query that can return multiple matches
func (ud *UserDialog) QueryAll(prompt string, validator *re.Regexp, key string, maxMatches int) (matches [][]string) {
	_ = "STUB: not implemented"
	return nil
}

// QueryString writes the prompt to the user and returns the answer if it passes the validator function
func (ud *UserDialog) QueryString(prompt string, validator func(string) error, key string) string {
	_ = "STUB: not implemented"
	return ""
}

// Input closed, so let's just return an empty string

// QueryStringPattern is a version of QueryString taking a regular expression pattern as the validator
func (ud *UserDialog) QueryStringPattern(prompt string, validator *re.Regexp, key string) (answer string) {
	_ = "STUB: not implemented"
	return ""
}

// QueryInt writes the prompt to the user and returns the answer if it can be parsed as an integer
func (ud *UserDialog) QueryInt(prompt string, key string, bitSize int) (value int64) {
	_ = "STUB: not implemented"
	return 0
}

// Explicitly treat #ffa080 as hexadecimal

// QueryBool writes the prompt to the user and returns the answer if it can be parsed as a boolean
func (ud *UserDialog) QueryBool(prompt string, key string) (value bool) {
	_ = "STUB: not implemented"
	return false
}
