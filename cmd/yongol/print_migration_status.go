//ff:func feature=cli type=util control=sequence
//ff:what printMigrationStatus — prints the Migration section of the yongol status dashboard
package main

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// printMigrationStatus writes a Migration Status section to w. It is
// defensive: any error is turned into a single status line so the
// dashboard never fails.
func printMigrationStatus(w io.Writer, specsDir, artsDir string) {
	fmt.Fprintln(w, "\n── Migration Status ──")
	ddlDir := filepath.Join(specsDir, migration.SnapshotSubdir)
	snapshotPath := filepath.Join(ddlDir, migration.SnapshotFileName)

	if !printMigrationSnapshotInfo(w, snapshotPath) {
		return
	}
	if !printMigrationPendingDiff(w, ddlDir, snapshotPath) {
		return
	}
	printMigrationLatestFile(w, artsDir)
}
