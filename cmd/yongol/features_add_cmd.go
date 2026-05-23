//ff:func feature=cli type=command control=sequence
//ff:what featuresAddCmd — yongol features add 서브커맨드

package main

import (
	"github.com/spf13/cobra"

	clifeatures "github.com/park-jun-woo/yongol/pkg/cmd/features"
)

// featuresAddCmd returns the `yongol features add <features.yaml>` subcommand.
func featuresAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "add <features.yaml>",
		Short:         "Add new features from a features.yaml file",
		Args:          usageArgs(cobra.ExactArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			specsDir := "specs"
			return clifeatures.RunAdd(cmd.OutOrStdout(), specsDir, args[0])
		},
	}
	return cmd
}
