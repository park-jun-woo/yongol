//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what D-2 test — NOT NULL 누락 시 ERROR / `-- @nullable` 면제 (BUG-028)

package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runD02InTmpDir writes sqlContent to <tmp>/db/<fname> and invokes the
// D-2 rule against a Fullstack whose SpecsDir points at <tmp>. Caller
// may pre-populate fs.DDLTables (e.g. to simulate parser-captured
// `-- @nullable` annotations) via the tables argument.
func runD02InTmpDir(t *testing.T, fname, sqlContent string, tables []ddl.Table) []string {
	t.Helper()
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, fname), []byte(sqlContent), 0o644); err != nil {
		t.Fatalf("write sql: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp, DDLTables: tables}
	diags := d02NullableColumn(fs)
	msgs := make([]string, len(diags))
	for i, d := range diags {
		msgs[i] = d.Message
	}
	return msgs
}

// TestD02NullableColumn_FlagsMissingNotNull — column without NOT NULL
// and without `-- @nullable` annotation must trigger [D-2].
func TestD02NullableColumn_FlagsMissingNotNull(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash  TEXT        PRIMARY KEY,
    claims      JSONB       NOT NULL,
    revoked_at  TIMESTAMPTZ
);`
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, nil)
	if len(msgs) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(msgs), msgs)
	}
	if !strings.Contains(msgs[0], "[D-2]") || !strings.Contains(msgs[0], "revoked_at") {
		t.Errorf("diag should mention [D-2] and revoked_at, got %q", msgs[0])
	}
}

// TestD02NullableColumn_RespectsNullableAnnotation — BUG-028 regression.
// When the parser has recorded `-- @nullable` on a column, D-2 must not
// fire for that column.
func TestD02NullableColumn_RespectsNullableAnnotation(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash  TEXT        PRIMARY KEY,
    claims      JSONB       NOT NULL,
    revoked_at  TIMESTAMPTZ -- @nullable
);`
	tables := []ddl.Table{{
		Name:          "refresh_tokens",
		NullableAnnot: map[string]bool{"revoked_at": true},
	}}
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, tables)
	if len(msgs) != 0 {
		t.Fatalf("expected 0 diagnostics (column is @nullable-annotated), got %d: %+v", len(msgs), msgs)
	}
}

// TestD02NullableColumn_PrimaryKeyExempt — PRIMARY KEY columns are
// implicitly NOT NULL and must not trigger [D-2] even without an
// explicit NOT NULL constraint.
func TestD02NullableColumn_PrimaryKeyExempt(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash TEXT PRIMARY KEY,
    claims     JSONB NOT NULL
);`
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, nil)
	if len(msgs) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(msgs), msgs)
	}
}
