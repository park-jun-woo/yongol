//ff:func feature=cli type=command control=sequence
//ff:what initCmd — yongol init <ProjectID> "<description>" subcommand
package main

import (
	"github.com/spf13/cobra"

	cliinit "github.com/park-jun-woo/yongol/pkg/cmd/init"
)

// initCmd returns the `yongol init` subcommand. The command materializes a
// minimal SSOT skeleton that `yongol validate specs` accepts with zero errors
// so new users can move straight into `yongol add` / `yongol get` without
// writing any boilerplate YAML by hand.
func initCmd() *cobra.Command {
	var (
		dirFlag    string
		moduleFlag string
		forceFlag  bool
	)
	cmd := &cobra.Command{
		Use:           `init <ProjectID> "<description>"`,
		Short:         "Scaffold a new yongol project (manifest + OpenAPI + sqlc + rego skeleton)",
		Args:          usageArgs(cobra.ExactArgs(2)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := cliinit.Options{
				ProjectID:   args[0],
				Description: args[1],
				Dir:         dirFlag,
				Module:      moduleFlag,
				Force:       forceFlag,
			}
			return cliinit.Run(cmd.OutOrStdout(), cmd.ErrOrStderr(), opts)
		},
	}
	cmd.Flags().StringVar(&dirFlag, "dir", "", "target directory (default: ./<ProjectID>)")
	cmd.Flags().StringVar(&moduleFlag, "module", "", "Go module path for manifest.backend.module (default: auto-detected)")
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "allow writing into a non-empty directory")
	return cmd
}
