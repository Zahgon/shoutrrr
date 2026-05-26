package jsonclient

// Error contains additional http/JSON details
type Error struct {
	StatusCode int
	Body       string
	err        error
}

func (je Error) Error() string { _ = "STUB: not implemented"; return "" }

func (je Error) String() string { _ = "STUB: not implemented"; return "" }

// ErrorBody returns the request body from an Error
func ErrorBody(e error) string { _ = "STUB: not implemented"; return "" }
