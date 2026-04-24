//ff:func feature=migration type=util control=sequence
//ff:what removeLegacyBaseline — pre-Phase010 위치의 .generated_schema.sql 을 발견하면 삭제 (BUG-034)
package migration

import (
	"log"
	"os"
	"path/filepath"
)

// removeLegacyBaseline deletes the pre-Phase010 baseline file at
// <ddlDir>/.generated_schema.sql if it still exists. Phase010 (BUG-034)
// moved the baseline to <artsDir>/db/.latest_schema.sql; leaving a
// stale copy under specs/db/ causes drift confusion and migration
// regeneration bugs. yongol is pre-release so no backward compatibility
// is kept — the old file is removed, with a one-line warning to the
// log. Callers do not need to handle the error: the helper is
// best-effort and never blocks generate.
func removeLegacyBaseline(ddlDir string) {
	oldPath := filepath.Join(ddlDir, LegacySnapshotFileName)
	if _, err := os.Stat(oldPath); err != nil {
		return
	}
	if err := os.Remove(oldPath); err != nil {
		log.Printf("yongol: failed to remove stale baseline at %s: %v (see BUG-034)", oldPath, err)
		return
	}
	log.Printf("yongol: removed stale baseline at %s (see BUG-034 / Phase010)", oldPath)
}
