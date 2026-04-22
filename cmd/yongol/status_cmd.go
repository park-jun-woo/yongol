//ff:func feature=cli type=command control=sequence
//ff:what statusCmd — returns the yongol status subcommand (SSOT Summary + Artifacts + Preserved + Drift)

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/contract"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/validate"
	vcontract "github.com/park-jun-woo/yongol/pkg/validate/contract"
)

// statusCmd wires `yongol status <specs-dir> [<arts-dir>]`. The command
// parses every SSOT, prints an SSOT summary, and when arts-dir is
// supplied extends the report with Artifacts / Preserved / Drift
// sections. It never fails on drift — status is a read-only dashboard,
// not a gate.
func statusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "status <specs-dir> [<arts-dir>]",
		Short:         "Print SSOT summary and (optional) artifact drift dashboard",
		Args:          usageArgs(cobra.RangeArgs(1, 2)),
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
				printParseErrors(cmd.OutOrStdout(), fs.ParseDiagnostics)
				return fmt.Errorf("parse failed")
			}
			out := cmd.OutOrStdout()
			printSSOTSummary(out, fs)
			// Migration status — always shown when DDL is present.
			hasDDL := false
			for _, d := range detected {
				if d.Kind == yongol.KindDDL {
					hasDDL = true
					break
				}
			}
			if hasDDL {
				artsForMig := ""
				if len(args) >= 2 {
					artsForMig = args[1]
				}
				printMigrationStatus(out, specsDir, artsForMig)
			}
			if len(args) < 2 {
				return nil
			}
			artsDir := args[1]
			if _, err := os.Stat(artsDir); err != nil {
				fmt.Fprintf(out, "\n(no artifacts at %s)\n", artsDir)
				return nil
			}
			// Validate populates Ground so contract.Run can resolve SSOT
			// lookups; the synthetic "contract" step is inspected via the
			// dedicated Run call below for dashboard-friendly formatting.
			_ = validate.Validate(fs, validate.WithArtsDir(artsDir))
			preserved, _ := contract.CollectPreserved(artsDir)
			diags := vcontract.Run(fs, artsDir)
			printArtifactsSummary(out, artsDir, len(preserved), len(diags))
			printPreservedList(out, preserved)
			printDriftList(out, diags)
			return nil
		},
	}
	return cmd
}
