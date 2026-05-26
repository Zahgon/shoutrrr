//go:generate stringer -type=URLPart -trimprefix URL

package format

// URLPart is an indicator as to what part of an URL a field is serialized to
type URLPart int

// Suffix returns the separator between the URLPart and it's subsequent part
func (u URLPart) Suffix() rune { _ = "STUB: not implemented"; return 0 }

// indicator as to what part of an URL a field is serialized to
const (
	URLQuery URLPart = iota
	URLUser
	URLPassword
	URLHost
	URLPort
	URLPath
)

// ParseURLPart returns the URLPart that matches the supplied string
func ParseURLPart(s string) URLPart { _ = "STUB: not implemented"; return *new(URLPart) }

// ParseURLParts returns the URLParts that matches the supplied string
func ParseURLParts(s string) []URLPart { _ = "STUB: not implemented"; return nil }
