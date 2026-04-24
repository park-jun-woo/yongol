//ff:func feature=migration type=generator control=sequence
//ff:what emitMigration — 새 마이그레이션 SQL 파일 작성 + 스냅샷 갱신 후 Result 반환
package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// emitMigration writes the numbered migration file and the canonical
// snapshot. Separated from Generate so Q3 doesn't fire.
func emitMigration(specsDir, artifactsDir string, curr *Schema, ops []Operation, hints *Hints, mode Mode, yongolVersion string, now time.Time, diags []diagnostic.Diagnostic) (*Result, []diagnostic.Diagnostic, error) {
	seq, err := NextSequenceNumber(filepath.Join(artifactsDir, MigrationsSubdir))
	if err != nil {
		return nil, diags, fmt.Errorf("next sequence: %w", err)
	}
	desc := InferDescription(ops)
	if mode == ModeInitial {
		desc = InitialMigrationDesc
	}
	upName := fmt.Sprintf(MigrationFilenameFormat, seq, desc)
	downName := fmt.Sprintf(MigrationDownFilenameFormat, seq, desc)
	upSQL := EmitSQL(ops, EmitOptions{YongolVersion: yongolVersion, GeneratedAt: now})
	downSQL := RenderDownStub(yongolVersion, now)

	migDir := filepath.Join(artifactsDir, MigrationsSubdir)
	if err := os.MkdirAll(migDir, 0755); err != nil {
		return nil, diags, fmt.Errorf("mkdir migrations: %w", err)
	}
	if err := os.WriteFile(filepath.Join(migDir, upName), []byte(upSQL), 0644); err != nil {
		return nil, diags, fmt.Errorf("write migration: %w", err)
	}
	if err := os.WriteFile(filepath.Join(migDir, downName), []byte(downSQL), 0644); err != nil {
		return nil, diags, fmt.Errorf("write down stub: %w", err)
	}

	baselineDir := filepath.Join(artifactsDir, BaselineSubdir)
	if err := os.MkdirAll(baselineDir, 0755); err != nil {
		return nil, diags, fmt.Errorf("mkdir baseline: %w", err)
	}
	snapshotPath := filepath.Join(baselineDir, SnapshotFileName)
	if err := WriteSnapshot(snapshotPath, curr, yongolVersion, now); err != nil {
		return nil, diags, fmt.Errorf("write snapshot: %w", err)
	}
	_ = specsDir // reserved: legacy baseline cleanup happens in Generate.

	return &Result{
		Mode:          mode,
		MigrationFile: upName,
		SnapshotFile:  filepath.Join(BaselineSubdir, SnapshotFileName),
		OpsCount:      len(ops),
		Operations:    ops,
		Hints:         hints,
		Warnings:      diags,
	}, diags, nil
}
