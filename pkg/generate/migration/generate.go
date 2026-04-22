//ff:func feature=migration type=command control=sequence
//ff:what Generate — migration 파이프라인 엔트리포인트 (스냅샷 로드 → diff → emit → 스냅샷 갱신)
package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// Mode indicates which branch the generator took.
type Mode string

const (
	ModeInitial     Mode = "initial"
	ModeIncremental Mode = "incremental"
	ModeNoop        Mode = "noop"
)

// Result is what Generate returns to the caller.
type Result struct {
	Mode          Mode
	MigrationFile string // relative to artifactsDir, "" on noop
	SnapshotFile  string // relative to specsDir
	OpsCount      int
	Operations    []Operation
	Hints         *Hints
	Warnings      []diagnostic.Diagnostic // non-blocking MIG-004/005 etc
}

// Options tunes Generate.
type Options struct {
	YongolVersion string
	Now            time.Time // for deterministic tests; zero => time.Now()
}

// Generate runs the full migration step. It returns errors that carry
// MIG-* ERROR diagnostics; the caller (yongol generate) should surface
// those via the normal report pipeline.
func Generate(specsDir, artifactsDir string, opt Options) (*Result, []diagnostic.Diagnostic, error) {
	ddlDir := filepath.Join(specsDir, SnapshotSubdir)
	snapshotPath := filepath.Join(ddlDir, SnapshotFileName)

	// 1. Load previous snapshot (if any).
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
		// Guard against "migrations exist but snapshot doesn't" per Phase001 Q2.
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

	// 2. Build current schema from DDL files (skip snapshot file itself).
	curr, err := BuildASTFromDir(ddlDir, []string{SnapshotFileName})
	if err != nil {
		return nil, errDiags, fmt.Errorf("build current AST: %w", err)
	}

	// 3. Collect hint comments + ParseHints.
	rawHints, err := ddl.ExtractHintCommentsFromDir(ddlDir)
	if err != nil {
		return nil, errDiags, fmt.Errorf("extract hint comments: %w", err)
	}
	hints := ParseHints(rawHints)

	// 4. Diff.
	ops := Diff(prev, curr, hints)
	ops = ApplyHintsToOps(ops, hints)

	// 5. Safety check + aggregate MIG-* diagnostics.
	issues := CheckSafety(ops)
	_, missing := LoadDataMigrationSQL(specsDir, hints)

	// Delegate to validate/migration.Run for MIG-* diagnostics.
	// Avoid import cycle by constructing diagnostics inline here; the
	// validate package is a consumer with the same data, not a producer
	// that Generate needs to call. (Phase004 rules fire from the
	// generate CLI hook too.)
	var diags []diagnostic.Diagnostic
	diags = append(diags, errDiags...)
	diags = append(diags, mig001From(prev, curr, hints)...)
	diags = append(diags, issuesToDiags(issues, missing)...)
	// MIG-006 runs on the snapshot file (already loaded cleanly at this point).

	// ERROR 하나라도 있으면 빠진다
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			return &Result{Mode: mode, Hints: hints, Operations: ops}, diags, fmt.Errorf("migration blocked by %d error(s)", countErrors(diags))
		}
	}

	// 6. No changes path.
	if len(ops) == 0 {
		return &Result{Mode: ModeNoop, Hints: hints}, diags, nil
	}

	// 7. Emit migration file.
	now := opt.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	seq, err := NextSequenceNumber(filepath.Join(artifactsDir, MigrationsSubdir))
	if err != nil {
		return nil, diags, fmt.Errorf("next sequence: %w", err)
	}
	desc := InferDescription(ops)
	if mode == ModeInitial {
		desc = InitialMigrationDesc
	}
	filename := fmt.Sprintf(MigrationFilenameFormat, seq, desc)
	sql := EmitSQL(ops, EmitOptions{YongolVersion: opt.YongolVersion, GeneratedAt: now})

	migPath := filepath.Join(artifactsDir, MigrationsSubdir, filename)
	if err := os.MkdirAll(filepath.Dir(migPath), 0755); err != nil {
		return nil, diags, fmt.Errorf("mkdir migrations: %w", err)
	}
	if err := os.WriteFile(migPath, []byte(sql), 0644); err != nil {
		return nil, diags, fmt.Errorf("write migration: %w", err)
	}

	// 8. Snapshot update (canonical SQL + hash header).
	if err := WriteSnapshot(snapshotPath, curr, opt.YongolVersion, now); err != nil {
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

func countErrors(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			n++
		}
	}
	return n
}
