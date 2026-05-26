package docs

import (
	"strings"

	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/spf13/cobra"

	cli "github.com/containrrr/shoutrrr/shoutrrr/cmd"
)

var (
	serviceRouter router.ServiceRouter
	services      = serviceRouter.ListServices()
)

// Cmd prints documentation for services
var Cmd = &cobra.Command{
	Use:   "docs",
	Short: "Print documentation for services",
	Run:   Run,
	Args: func(cmd *cobra.Command, args []string) error {
		serviceList := strings.Join(services, ", ")
		cmd.SetUsageTemplate(cmd.UsageTemplate() + "\nAvailable services: \n  " + serviceList + "\n")
		return cobra.MinimumNArgs(1)(cmd, args)
	},
	ValidArgs: services,
}

func init() {
	Cmd.Flags().StringP("format", "f", "console", "Output format")
}

// Run the docs command
func Run(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

func printDocs(format string, services []string) cli.Result {
	_ = "STUB: not implemented"
	return *new(cli.Result)
}
