package standard

import (
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Logger provides the utility methods Log* that maps to Logger.Print*
type Logger struct {
	logger types.StdLogger
}

// Logf maps to the service loggers Logger.Printf function
func (sl *Logger) Logf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Log maps to the service loggers Logger.Print function
func (sl *Logger) Log(v ...interface{}) { _ = "STUB: not implemented"; return }

// SetLogger maps the specified logger to the Log* helper methods
func (sl *Logger) SetLogger(logger types.StdLogger) { _ = "STUB: not implemented"; return }
