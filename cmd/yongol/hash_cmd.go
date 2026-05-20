//ff:func feature=cli type=command control=sequence
//ff:what hashCmd — yongol hash <specs-dir> subcommand
package main

import (
	"github.com/spf13/cobra"

	clihash "github.com/park-jun-woo/yongol/pkg/cmd/hash"
)

// hashCmd returns the `yongol hash` subcommand. The command reads
// features.yaml inside the given specs directory, computes its SHA-256,
// and writes (or overwrites) the .yongol hash lock file.
func hashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "hash <specs-dir>",
		Short:         "Generate .yongol hash lock from features.yaml",
		Args:          usageArgs(cobra.ExactArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return clihash.Run(cmd.OutOrStdout(), args[0])
		},
	}
	return cmd
}
