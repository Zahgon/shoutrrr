package testutils

import (
	"io"
)

type failWriter struct {
	writeLimit int
	writeCount int
}

// Close is just a dummy function to implement io.Closer
func (fw *failWriter) Close() error {
	_ = "STUB: not implemented"

	// Write returns an error if the write limit has been reached
	return nil
}

func (fw *failWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CreateFailWriter returns a io.WriteCloser that returns an error after the amount of writes indicated by writeLimit
func CreateFailWriter(writeLimit int) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}
