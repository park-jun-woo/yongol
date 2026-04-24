//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD09 — top-level INSERT lacking `-- @sentinel` raises ERROR
package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD09_UnannotatedInsertFlagged(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir db: %v", err)
	}
	sql := `CREATE TABLE t (id BIGINT);

INSERT INTO t (id) VALUES (0) ON CONFLICT DO NOTHING;
`
	if err := os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644); err != nil {
		t.Fatalf("write sql: %v", err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d09TopLevelInsertWithoutSentinel(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[D-9]") {
		t.Errorf("expected [D-9] prefix: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Advice, "-- @sentinel") {
		t.Errorf("advice should mention -- @sentinel: %q", diags[0].Advice)
	}
}

func TestD09_AnnotatedInsertOK(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);

-- @sentinel
INSERT INTO t (id) VALUES (0) ON CONFLICT DO NOTHING;
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d09TopLevelInsertWithoutSentinel(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}

func TestD09_NoInsertOK(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	_ = os.MkdirAll(dbDir, 0o755)
	sql := `CREATE TABLE t (id BIGINT);
`
	_ = os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644)
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d09TopLevelInsertWithoutSentinel(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
