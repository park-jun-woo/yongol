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
	filename := fmt.Sprintf(MigrationFilenameFormat, seq, desc)
	sql := EmitSQL(ops, EmitOptions{YongolVersion: yongolVersion, GeneratedAt: now})

	migPath := filepath.Join(artifactsDir, MigrationsSubdir, filename)
	if err := os.MkdirAll(filepath.Dir(migPath), 0755); err != nil {
		return nil, diags, fmt.Errorf("mkdir migrations: %w", err)
	}
	if err := os.WriteFile(migPath, []byte(sql), 0644); err != nil {
		return nil, diags, fmt.Errorf("write migration: %w", err)
	}

	snapshotPath := filepath.Join(specsDir, SnapshotSubdir, SnapshotFileName)
	if err := WriteSnapshot(snapshotPath, curr, yongolVersion, now); err != nil {
		return nil, diags, fmt.Errorf("write snapshot: %w", err)
	}

	return &Result{
		Mode:          mode,
		MigrationFile: filename,
		SnapshotFile:  filepath.Join(SnapshotSubdir, SnapshotFileName),
		OpsCount:      len(ops),
		Operations:    ops,
		Hints:         hints,
		Warnings:      diags,
	}, diags, nil
}
