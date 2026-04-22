//ff:func feature=cli type=command control=sequence
//ff:what chainCmd — returns the yongol chain subcommand
package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/chain"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func chainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "chain <operationId> <specs-dir>",
		Short:         "Trace all SSOT nodes connected to an operationId",
		Args:          usageArgs(cobra.ExactArgs(2)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			opID, specsDir := args[0], args[1]
			detected, err := yongol.DetectSSOTs(specsDir)
			if err != nil {
				return fmt.Errorf("detect SSOTs: %w", err)
			}
			fs := yongol.ParseAll(specsDir, detected)
			if len(fs.ParseDiagnostics) > 0 {
				printParseErrors(cmd.OutOrStdout(), fs.ParseDiagnostics)
				return fmt.Errorf("parse failed")
			}
			links, err := chain.Chain(fs, opID)
			if err != nil {
				return err
			}
			chain.Print(cmd.OutOrStdout(), opID, links)
			return nil
		},
	}
	return cmd
}
