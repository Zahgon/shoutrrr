package types

// Params is the string map used to provide additional variables to the service templates
type Params map[string]string

const (
	// TitleKey is the common key for the title prop
	TitleKey = "title"
	// MessageKey is the common key for the message prop
	MessageKey = "message"
)

// SetTitle sets the "title" param to the specified value
func (p Params) SetTitle(title string) { _ = "STUB: not implemented"; return }

// Title returns the "title" param
func (p Params) Title() (title string, found bool) { _ = "STUB: not implemented"; return "", false }

// SetMessage sets the "message" param to the specified value
func (p Params) SetMessage(message string) { _ = "STUB: not implemented"; return }
