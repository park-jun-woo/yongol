//ff:func feature=cli type=command control=sequence
//ff:what nextCmd — returns the yongol next subcommand (groups errors by operationId for iterative fixing)
package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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

//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what collectIssues — Report에서 모든 ERROR/WARNING 수집 (ERROR 우선 정렬)
func collectIssues(r *validate.Report) []diagnostic.Diagnostic {
	var errors, warnings []diagnostic.Diagnostic
	for _, s := range r.Steps {
		for _, d := range s.Diagnostics {
			switch d.Level {
			case diagnostic.LevelError:
				errors = append(errors, d)
			case diagnostic.LevelWarning:
				warnings = append(warnings, d)
			}
		}
	}
	return append(errors, warnings...)
}

//ff:func feature=cli type=helper control=sequence
//ff:what printGroupedDiags — 같은 OperationID의 진단을 묶어서 출력
func printGroupedDiags(out io.Writer, opID string, all []diagnostic.Diagnostic, specsDir string) {
	var group []diagnostic.Diagnostic
	for _, d := range all {
		if d.OperationID == opID {
			group = append(group, d)
		}
	}
	fmt.Fprintf(out, "\n%s (%d errors):\n", opID, len(group))
	for _, d := range group {
		fmt.Fprintf(out, "  [%s] %s\n", d.Level, d.Message)
		if d.File != "" {
			if d.Line > 0 {
				fmt.Fprintf(out, "    file: %s:%d\n", d.File, d.Line)
			} else {
				fmt.Fprintf(out, "    file: %s\n", d.File)
			}
		}
	}
	fmt.Fprintf(out, "  ▶ Fix all errors for %s. Then run `yongol next %s`. Don't stop. Keep going until complete.\n", opID, specsDir)
}

//ff:func feature=cli type=helper control=sequence
//ff:what printNextDiag — 단일 diagnostic을 next 형식으로 출력 (1줄 에러 + file + fix 지시)
func printNextDiag(cmd *cobra.Command, d diagnostic.Diagnostic, specsDir string) {
	out := cmd.OutOrStdout()
	fmt.Fprintf(out, "\n[%s] %s\n", d.Level, d.Message)
	if d.File != "" {
		if d.Line > 0 {
			fmt.Fprintf(out, "  file: %s:%d\n", d.File, d.Line)
		} else {
			fmt.Fprintf(out, "  file: %s\n", d.File)
		}
	}
	fmt.Fprintf(out, "  ▶ Fix this error. Then run `yongol next %s`. Don't stop. Keep going until complete.\n", specsDir)
}
