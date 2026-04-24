//ff:func feature=validate type=test control=iteration dimension=2 topic=query-structural
//ff:what Q-11 — sqlc.yaml sql_package 가 pgx/v5 아닐 때 ERROR 발화

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// writeSqlcYaml drops a minimal sqlc.yaml with the given sql_package into
// <dir>/db/sqlc.yaml. An empty pkg writes the field absent entirely.
func writeSqlcYaml(t *testing.T, dir string, pkg string) {
	t.Helper()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	var body string
	if pkg == "" {
		body = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
`
	} else {
		body = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
        sql_package: "` + pkg + `"
`
	}
	if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(body), 0o644); err != nil {
		t.Fatalf("write sqlc.yaml: %v", err)
	}
}

func TestQ11SqlPackagePgxV5_Pass(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "pgx/v5")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) != 0 {
		t.Fatalf("pgx/v5 must pass Q-11, got %d diagnostics: %+v", len(diags), diags)
	}
}

func TestQ11SqlPackagePgxV5_DatabaseSqlFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "database/sql")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("database/sql must fire Q-11, got no diagnostics")
	}
	if !strings.Contains(diags[0].Message, "[Q-11]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, `"database/sql"`) {
		t.Errorf("current value missing from message: %q", diags[0].Message)
	}
}

func TestQ11SqlPackagePgxV5_PgxV4Fails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "pgx/v4")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("pgx/v4 must fire Q-11, got no diagnostics")
	}
}

func TestQ11SqlPackagePgxV5_AbsentFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("absent sql_package must fire Q-11, got no diagnostics")
	}
	if !strings.Contains(diags[0].Message, "absent") {
		t.Errorf("message should note absence: %q", diags[0].Message)
	}
}

func TestQ11SqlPackagePgxV5_LibPqFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "lib/pq")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("lib/pq must fire Q-11, got no diagnostics")
	}
}
