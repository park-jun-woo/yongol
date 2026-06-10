//ff:func feature=validate type=rule control=sequence topic=migration-snapshot
//ff:what MIG-006 — arts/db/.latest_schema.sql 의 YONGOL_SCHEMA_HASH 헤더가 본문 sha256 과 불일치 → ERROR
package migration

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Mig006SnapshotDrift verifies the snapshot file's header hash matches
// the sha256 of the remaining canonical SQL body. A drift implies the
// user (or AI) hand-edited the snapshot, which invalidates diff.
//
// As of Phase010 (BUG-034) the snapshot lives at
// <artsDir>/db/.latest_schema.sql — directly under the db directory so
// external migration tools (golang-migrate / flyway / goose) do not
// mistake it for a migration. If the file is absent the rule is a
// no-op (first-run initial is valid).
func Mig006SnapshotDrift(artsDir string) []diagnostic.Diagnostic {
	snap := filepath.Join(artsDir, migration.BaselineSubdir, migration.SnapshotFileName)
	data, err := os.ReadFile(snap)
	if err != nil {
		return nil // absent = initial mode; handled by generate entrypoint.
	}
	text := string(data)
	headerLine, body, ok := splitHashHeader(text)
	if !ok {
		return []diagnostic.Diagnostic{{
			File:    snap,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-006] snapshot has no YONGOL_SCHEMA_HASH header (drift)",
			Advice:  "regenerate by `yongol generate` — do not hand-edit the snapshot",
		}}
	}
	stored := strings.TrimSpace(strings.TrimPrefix(headerLine, migration.SnapshotHashHeaderPrefix))
	sum := sha256.Sum256([]byte(body))
	want := hex.EncodeToString(sum[:])
	if stored != want {
		return []diagnostic.Diagnostic{{
			File:    snap,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-006] snapshot drift — YONGOL_SCHEMA_HASH does not match body",
			Advice:  "run `yongol generate` to regenerate the snapshot; do not hand-edit",
		}}
	}
	return nil
}
