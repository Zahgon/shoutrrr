package util

import (
	"github.com/spf13/cobra"
)

// LoadFlagsFromAltSources is a WORKAROUND to make cobra count env vars and positional arguments when checking required flags
func LoadFlagsFromAltSources(cmd *cobra.Command, args []string) { _ = "STUB: not implemented"; return }

// If the URL has been set in ENV, default the message to read from stdin

func hasURLInEnvButNotFlag(cmd *cobra.Command) bool { _ = "STUB: not implemented"; return false }
