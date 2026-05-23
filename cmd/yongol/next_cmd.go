//ff:func feature=cli type=command control=sequence
//ff:what nextCmd — returns the yongol next subcommand (groups errors by operationId for iterative fixing)
package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/validate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// nextCmd wires `yongol next <specs-dir>`. The command runs the same
// validation logic as `yongol validate` but outputs errors grouped by
// operationId. When the first issue has an OperationID, all issues
// sharing that OperationID are printed together so the agent can fix
// related cross-layer errors in one pass.
func nextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "next <specs-dir>",
		Short:         "Print the first error group by operationId (iterative fixing loop)",
		Args:          usageArgs(cobra.ExactArgs(1)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			specsDir := args[0]
			detected, err := yongol.DetectSSOTs(specsDir)
			if err != nil {
				return fmt.Errorf("detect SSOTs: %w", err)
			}
			fs := yongol.ParseAll(specsDir, detected)
			if len(fs.ParseDiagnostics) > 0 {
				d := fs.ParseDiagnostics[0]
				printNextDiag(cmd, d, specsDir)
				return fmt.Errorf("parse failed")
			}
			report := validate.Validate(fs)
			issues := collectIssues(report)
			if len(issues) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), "✓ All validations passed. 0 errors.")
				return nil
			}
			first := issues[0]
			if first.OperationID != "" {
				printGroupedDiags(cmd.OutOrStdout(), first.OperationID, issues, specsDir)
			} else {
				printNextDiag(cmd, first, specsDir)
			}
			return fmt.Errorf("validation failed")
		},
	}
	return cmd
}
