package send

import (
	"github.com/spf13/cobra"

	intutil "github.com/containrrr/shoutrrr/internal/util"
)

// Cmd sends a notification using a service url
var Cmd = &cobra.Command{
	Use:    "send",
	Short:  "Send a notification using a service url",
	Args:   cobra.MaximumNArgs(2),
	PreRun: intutil.LoadFlagsFromAltSources,
	RunE:   Run,
}

func init() {
	Cmd.Flags().BoolP("verbose", "v", false, "")

	Cmd.Flags().StringArrayP("url", "u", []string{}, "The notification url")
	_ = Cmd.MarkFlagRequired("url")

	Cmd.Flags().StringP("message", "m", "", "The message to send to the notification url, or - to read message from stdin")
	_ = Cmd.MarkFlagRequired("message")

	Cmd.Flags().StringP("title", "t", "", "The title used for services that support it")
}

func logf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

func run(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// Only display "URLs:" prefix for first line, replace with indentation for the the subsequent

// Run the send command
func Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// If the error is not related to the CLI usage, report error and exit to not invoke cobra error output
