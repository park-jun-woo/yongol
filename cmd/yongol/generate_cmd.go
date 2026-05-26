//ff:func feature=cli type=command control=sequence
//ff:what generateCmd — returns the yongol generate subcommand

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/generate"
	"github.com/park-jun-woo/yongol/pkg/validate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateCmd wires `yongol generate <specs-dir> <artifacts-dir>`. Flow:
//  1. DetectSSOTs + ParseAll (identical to validate)
//  2. Fail on any parser diagnostic
//  3. Run the full validate pipeline
//  4. Fail when the validate report has any ERROR OR WARNING
//     (stricter than validate which only fails on ERRORs)
//  5. Call generate.Generate for the chosen backend / frontend targets
func generateCmd() *cobra.Command {
	var (
		backendFlag  string
		frontendFlag string
	)
	cmd := &cobra.Command{
		Use:           "generate <specs-dir> <artifacts-dir>",
		Short:         "Generate code artifacts from SSOTs",
		Args:          usageArgs(cobra.ExactArgs(2)),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			specsDir := args[0]
			artifactsDir := args[1]
			detected, err := yongol.DetectSSOTs(specsDir)
			if err != nil {
				return fmt.Errorf("detect SSOTs: %w", err)
			}
			fs := yongol.ParseAll(specsDir, detected)
			if len(fs.ParseDiagnostics) > 0 {
				printParseErrors(cmd.OutOrStdout(), fs.ParseDiagnostics)
				return fmt.Errorf("parse failed")
			}
			report := validate.Validate(fs)
			_, warns, err := printReport(cmd.OutOrStdout(), report, formatMD, specsDir)
			if err != nil {
				return err
			}
			if warns > 0 {
				return fmt.Errorf("generate refused: %d warnings must be resolved first", warns)
			}
			backend := generate.BackendType(backendFlag)
			switch backend {
			case generate.GoGin, generate.NestJS, generate.FastAPI:
				// valid
			default:
				if fs.Manifest != nil && fs.Manifest.Backend.Lang != "" {
					var err error
					backend, err = generate.ResolveBackendType(fs.Manifest.Backend.Lang, fs.Manifest.Backend.Framework)
					if err != nil {
						return fmt.Errorf("resolve backend from manifest: %w", err)
					}
				} else {
					return fmt.Errorf("unknown --backend value %q; valid: go-gin, nestjs, fastapi", backendFlag)
				}
			}
			frontend := generate.FrontendType(frontendFlag)
			migHook := generate.WithMigration(generate.MigrationHook{
				Version: Version,
				Logger:  cmd.OutOrStdout(),
			})
			if err := generate.Generate(fs, artifactsDir, backend, frontend, migHook); err != nil {
				return fmt.Errorf("generate: %w", err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "\nartifacts written to %s (backend=%s, frontend=%s)\n",
				artifactsDir, backend, frontend)
			return nil
		},
	}
	cmd.Flags().StringVar(&backendFlag, "backend", string(generate.GoGin), "backend framework: go-gin (default), nestjs, fastapi")
	cmd.Flags().StringVar(&frontendFlag, "frontend", string(generate.React), "frontend code generator (react)")
	return cmd
}
