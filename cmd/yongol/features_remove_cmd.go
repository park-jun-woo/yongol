//ff:func feature=cli type=command control=sequence
//ff:what featuresRemoveCmd — yongol features remove 서브커맨드

package main

import (
	"github.com/spf13/cobra"

	clifeatures "github.com/park-jun-woo/yongol/pkg/cmd/features"
)

// featuresRemoveCmd returns the `yongol features remove <operationId> [...]` subcommand.
func featuresRemoveCmd() *cobra.Command {
	var yesFlag bool
	cmd := &cobra.Command{
		Use:           "remove <operationId> [operationId...]",
		Short:         "Remove features by operationId",
		Args:          usageArgs(cobra.MinimumNArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			specsDir := "specs"
			return clifeatures.RunRemove(cmd.OutOrStdout(), cmd.InOrStdin(), specsDir, args, yesFlag)
		},
	}
	cmd.Flags().BoolVar(&yesFlag, "yes", false, "skip confirmation prompt")
	return cmd
}
