//ff:func feature=cli type=command control=sequence
//ff:what nextCmd — returns the yongol next subcommand (prints only the first validation error for iterative fixing)
package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// nextCmd wires `yongol next <specs-dir>`. The command runs the same
// validation logic as `yongol validate` but outputs only the FIRST error
// with a prompt to fix it and re-run. This one-at-a-time pattern helps
// agents converge by eliminating multi-error confusion.
func nextCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "next <specs-dir>",
		Short:         "Print only the first validation error (iterative fixing loop)",
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
			first, ok := firstError(report)
			if !ok {
				fmt.Fprintln(cmd.OutOrStdout(), "✓ All validations passed. 0 errors.")
				return nil
			}
			printNextDiag(cmd, first, specsDir)
			return fmt.Errorf("validation failed")
		},
	}
	return cmd
}

//ff:func feature=cli type=helper control=selection dimension=1
//ff:what firstError — Report에서 첫 번째 ERROR 레벨 diagnostic 반환
func firstError(r *validate.Report) (diagnostic.Diagnostic, bool) {
	for _, s := range r.Steps {
		for _, d := range s.Diagnostics {
			if d.Level == diagnostic.LevelError {
				return d, true
			}
		}
	}
	return diagnostic.Diagnostic{}, false
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
	fmt.Fprintf(out, "  ▶ Fix this error. Then run `yongol next %s`.\n", specsDir)
}
