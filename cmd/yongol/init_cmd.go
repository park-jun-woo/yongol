//ff:func feature=cli type=command control=sequence
//ff:what initCmd — yongol init <ProjectID> <features.yaml> ["description"] subcommand
package main

import (
	"github.com/spf13/cobra"

	cliinit "github.com/park-jun-woo/yongol/pkg/cmd/init"
)

// initCmd returns the `yongol init` subcommand. The command reads a
// features.yaml file and materializes SSOT stubs (OpenAPI, SSaC, Rego, Hurl)
// plus a hash lock (specs/.yongol) so the project is ready for iterative
// feature implementation.
func initCmd() *cobra.Command {
	var (
		dirFlag    string
		moduleFlag string
		forceFlag  bool
	)
	cmd := &cobra.Command{
		Use:           `init <ProjectID> <features.yaml> ["description"]`,
		Short:         "Scaffold a new yongol project from features.yaml",
		Args:          usageArgs(cobra.RangeArgs(2, 3)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var desc string
			if len(args) >= 3 {
				desc = args[2]
			}
			opts := cliinit.Options{
				ProjectID:    args[0],
				FeaturesPath: args[1],
				Description:  desc,
				Dir:          dirFlag,
				Module:       moduleFlag,
				Force:        forceFlag,
			}
			return cliinit.Run(cmd.OutOrStdout(), cmd.ErrOrStderr(), opts)
		},
	}
	cmd.Flags().StringVar(&dirFlag, "dir", "", "target directory (default: ./<ProjectID>)")
	cmd.Flags().StringVar(&moduleFlag, "module", "", "Go module path for manifest.backend.module (default: auto-detected)")
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "allow writing into a non-empty directory")
	return cmd
}
