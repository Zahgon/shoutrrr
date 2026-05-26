package failures

// FailureID is a number to be used to identify a specific error
type FailureID int

type failure struct {
	message string
	id      FailureID
	wrapped error
}

// Failure is an extended error that also includes an ID to be used to identify a specific error
type Failure interface {
	error
	ID() FailureID
}

func (f *failure) Error() string { _ = "STUB: not implemented"; return "" }

func (f *failure) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (f *failure) ID() FailureID { _ = "STUB: not implemented"; return *new(FailureID) }

func (f *failure) Is(target error) bool { _ = "STUB: not implemented"; return false }

// Wrap returns a failure with the given message and id, saving the message of wrappedError for appending to Error()
func Wrap(message string, id FailureID, wrappedError error, v ...interface{}) Failure {
	_ = "STUB: not implemented"
	return *new(Failure)
}

var _ error = &failure{}
