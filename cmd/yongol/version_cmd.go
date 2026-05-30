//ff:func feature=cli type=command control=sequence
//ff:what versionCmd — returns the yongol version subcommand
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print yongol version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "yongol %s\n", Version)
			fmt.Fprintf(cmd.OutOrStdout(), "  must read before editing SSOT: %s\n", manualForAIPath())
		},
	}
}
