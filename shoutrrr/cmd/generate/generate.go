package generate

import (
	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/spf13/cobra"
)

var serviceRouter router.ServiceRouter

// Cmd is used to generate a notification service URL from user input
var Cmd = &cobra.Command{
	Use:    "generate",
	Short:  "Generates a notification service URL from user input",
	Run:    Run,
	PreRun: loadArgsFromAltSources,
	Args:   cobra.MaximumNArgs(2),
}

func loadArgsFromAltSources(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func init() {
	serviceRouter = router.ServiceRouter{}
	Cmd.Flags().StringP("service", "s", "", "The notification service to generate a URL for")

	Cmd.Flags().StringP("generator", "g", "basic", "The generator to use")

	Cmd.Flags().StringArrayP("property", "p", []string{}, "Configuration property in key=value format")
}

// Run the generate command
func Run(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// try to use the service default generator if one exists
