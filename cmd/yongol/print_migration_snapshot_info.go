//ff:func feature=cli type=util control=sequence
//ff:what printMigrationSnapshotInfo — snapshot 존재/해시/DRIFT 상태 한 줄 출력 (성공 여부 반환)
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// printMigrationSnapshotInfo writes the snapshot status line. Returns false
// when a non-os.IsNotExist read error occurred so the caller stops the
// Migration section (the pending diff + latest file lines would be
// misleading).
func printMigrationSnapshotInfo(w io.Writer, snapshotPath string) bool {
	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(w, "snapshot: (absent — next generate will emit 0001_initial.up.sql)\n")
			return true
		}
		fmt.Fprintf(w, "snapshot: ERROR %s\n", err)
		return false
	}
	hash, drift := snapshotHashInfo(data)
	short := hash
	if len(short) > 8 {
		short = short[:8]
	}
	mark := "ok"
	if drift {
		mark = "DRIFT (MIG-006)"
	}
	fmt.Fprintf(w, "snapshot: %s  hash=%s  %s\n",
		filepath.Join(migration.BaselineSubdir, migration.SnapshotFileName), short, mark)
	return true
}
