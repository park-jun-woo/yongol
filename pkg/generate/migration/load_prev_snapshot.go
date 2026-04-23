//ff:func feature=migration type=loader control=sequence
//ff:what loadPrevSnapshot — 이전 스냅샷 로드 + 초기 모드 판정 + state inconsistency MIG-006 진단
package migration

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// loadPrevSnapshot loads the previous schema snapshot and decides the
// mode (initial / incremental). It also reports MIG-006 diagnostics when
// the snapshot fails to load or when the project is in an inconsistent
// state (migrations exist but snapshot doesn't).
func loadPrevSnapshot(snapshotPath, artifactsDir string) (*Schema, Mode, []diagnostic.Diagnostic) {
	prev, prevErr := LoadSnapshot(snapshotPath)
	var errDiags []diagnostic.Diagnostic
	if prevErr != nil {
		errDiags = append(errDiags, diagnostic.Diagnostic{
			File:    snapshotPath,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-006] snapshot load failed: " + prevErr.Error(),
			Advice:  "regenerate by `yongol generate` or delete the file for a fresh initial",
		})
	}
	mode := ModeIncremental
	if prev == nil {
		mode = ModeInitial
		prev = NewSchema()
		migDir := filepath.Join(artifactsDir, MigrationsSubdir)
		if entries, _ := os.ReadDir(migDir); len(entries) > 0 {
			errDiags = append(errDiags, diagnostic.Diagnostic{
				File:    migDir,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[MIG-006] migrations exist but snapshot is absent (state inconsistent)",
				Advice:  "remove artifacts/db/migrations to start fresh, or restore the snapshot from git",
			})
		}
	}
	return prev, mode, errDiags
}
