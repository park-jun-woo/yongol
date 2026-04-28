//ff:func feature=validate type=test-helper control=sequence topic=query-structural
//ff:what runQ12PgtypeUuidOverrideCase — 단일 Q-12 케이스 (DDL+sqlc.yaml 작성, 진단 비교)

package query

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runQ12PgtypeUuidOverrideCase executes one row of the Q-12 table:
// writes DDL + sqlc.yaml into a tempdir, runs the rule, then delegates
// the assertion to assertQ12Diags. Extracted from TestQ12PgtypeUuidOverride
// to keep the test loop body within the Q4 PURE line budget.
func runQ12PgtypeUuidOverrideCase(t *testing.T, tc q12UuidTestCase) {
	t.Helper()
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte(tc.ddl), 0o644); err != nil {
		t.Fatalf("write users.sql: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(tc.sqlc), 0o644); err != nil {
		t.Fatalf("write sqlc.yaml: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q12PgtypeUuidOverride(fs)
	assertQ12Diags(t, tc, diags)
}
