package standard

import (
	f "github.com/containrrr/shoutrrr/internal/failures"
)

const (
	// FailTestSetup is the FailureID used to represent an error that is part of the setup for tests
	FailTestSetup f.FailureID = -1
	// FailParseURL is the FailureID used to represent failing to parse the service URL
	FailParseURL f.FailureID = -2
	// FailServiceInit is the FailureID used to represent failure of a service.Initialize method
	FailServiceInit f.FailureID = -3
	// FailUnknown is the default FailureID
	FailUnknown f.FailureID = iota
)

// Failure creates a Failure instance corresponding to the provided failureID, wrapping the provided error
func Failure(failureID f.FailureID, err error, v ...interface{}) f.Failure {
	_ = "STUB: not implemented"
	return *new(f.Failure)
}

type failureLike interface {
	f.Failure
}

// IsTestSetupFailure checks whether the given failure is due to the test setup being broken
func IsTestSetupFailure(failure failureLike) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
