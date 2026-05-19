//ff:func feature=cli type=command control=sequence
//ff:what featuresCmd — yongol features add/remove 서브커맨드

package main

import (
	"github.com/spf13/cobra"

	clifeatures "github.com/park-jun-woo/yongol/pkg/cmd/features"
)

// featuresCmd returns the `yongol features` parent command with add/remove
// subcommands for managing features.yaml after init.
func featuresCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "features",
		Short: "Manage features.yaml (add/remove features)",
	}

	cmd.AddCommand(featuresAddCmd())
	cmd.AddCommand(featuresRemoveCmd())

	return cmd
}

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
