//ff:func feature=cli type=command control=sequence
//ff:what validateCmd — yongol validate 서브커맨드 반환 (--format md|sarif 포함)
package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func validateCmd() *cobra.Command {
	var formatFlag string
	cmd := &cobra.Command{
		Use:           "validate <specs-dir> [arts-dir]",
		Short:         "Validate SSOTs under specs-dir (optionally running contract drift checks against arts-dir)",
		Args:          usageArgs(cobra.RangeArgs(1, 2)),
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			switch formatFlag {
			case "", formatMD, formatJSON, formatSARIF:
				return nil
			default:
				return &usageError{err: fmt.Errorf("invalid --format %q (supported: md, json, sarif)", formatFlag)}
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			specsDir := args[0]
			var artsDir string
			if len(args) >= 2 {
				artsDir = args[1]
			}
			detected, err := yongol.DetectSSOTs(specsDir)
			if err != nil {
				return fmt.Errorf("detect SSOTs: %w", err)
			}
			fs := yongol.ParseAll(specsDir, detected)
			if len(fs.ParseDiagnostics) > 0 {
				printParseErrors(cmd.OutOrStdout(), fs.ParseDiagnostics)
				return fmt.Errorf("parse failed")
			}
			var opts []validate.Option
			if artsDir != "" {
				opts = append(opts, validate.WithArtsDir(artsDir))
			}
			report := validate.Validate(fs, opts...)
			_, _, err = printReport(cmd.OutOrStdout(), report, formatFlag, specsDir)
			return err
		},
	}
	cmd.Flags().StringVarP(&formatFlag, "format", "f", formatMD,
		"output format: md (GitHub Flavored Markdown, default) | json (flat snake_case) | sarif (SARIF 2.1.0 full catalog)")
	return cmd
}
