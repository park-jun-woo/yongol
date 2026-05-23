//ff:func feature=cli type=command control=sequence
//ff:what featuresCmd — yongol features add/remove 서브커맨드

package main

import (
	"github.com/spf13/cobra"
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
