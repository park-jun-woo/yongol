//ff:func feature=migration type=command control=sequence
//ff:what Generate — migration 파이프라인 엔트리포인트 (스냅샷 로드 → diff → emit → 스냅샷 갱신)
package migration

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// Generate runs the full migration step. It returns errors that carry
// MIG-* ERROR diagnostics; the caller (yongol generate) should surface
// those via the normal report pipeline.
func Generate(specsDir, artifactsDir string, opt Options) (*Result, []diagnostic.Diagnostic, error) {
	ddlDir := filepath.Join(specsDir, SnapshotSubdir)
	snapshotPath := filepath.Join(ddlDir, SnapshotFileName)

	prev, mode, errDiags := loadPrevSnapshot(snapshotPath, artifactsDir)

	curr, err := BuildASTFromDir(ddlDir, []string{SnapshotFileName})
	if err != nil {
		return nil, errDiags, fmt.Errorf("build current AST: %w", err)
	}

	rawHints, err := ddl.ExtractHintCommentsFromDir(ddlDir)
	if err != nil {
		return nil, errDiags, fmt.Errorf("extract hint comments: %w", err)
	}
	hints := ParseHints(rawHints)

	ops := Diff(prev, curr, hints)
	ops = ApplyHintsToOps(ops, hints)

	issues := CheckSafety(ops)
	_, missing := LoadDataMigrationSQL(specsDir, hints)

	diags := collectDiags(errDiags, prev, curr, hints, issues, missing)

	if hasErrorDiag(diags) {
		return &Result{Mode: mode, Hints: hints, Operations: ops}, diags,
			fmt.Errorf("migration blocked by %d error(s)", countErrors(diags))
	}

	if len(ops) == 0 {
		return &Result{Mode: ModeNoop, Hints: hints}, diags, nil
	}

	now := opt.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	return emitMigration(specsDir, artifactsDir, curr, ops, hints, mode, opt.YongolVersion, now, diags)
}
