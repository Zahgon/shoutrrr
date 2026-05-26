package cmd

// Result contains the final exit message and code for a CLI session
type Result struct {
	ExitCode int
	Message  string
}

func (e Result) Error() string {
	_ = "STUB: not implemented"

	// Success is the empty Result that is used whenever the command ran successfully
	return ""
}

var Success = Result{}

// InvalidUsage returns a Result with the exit code ExUsage
func InvalidUsage(message string) Result { _ = "STUB: not implemented"; return *new(Result) }

// TaskUnavailable returns a Result with the exit code ExUnavailable
func TaskUnavailable(message string) Result { _ = "STUB: not implemented"; return *new(Result) }

// ConfigurationError returns a Result with the exit code ExConfig
func ConfigurationError(message string) Result { _ = "STUB: not implemented"; return *new(Result) }

const (
	//ExSuccess is the exit code that signals that everything went as expected
	ExSuccess = 0
	//ExUsage is the exit code that signals that the application was not started with the correct arguments
	ExUsage = 64
	//ExUnavailable is the exit code that signals that the application failed to perform the intended task
	ExUnavailable = 69
	//ExConfig is the exit code that signals that the task failed due to a configuration error
	ExConfig = 78
)
