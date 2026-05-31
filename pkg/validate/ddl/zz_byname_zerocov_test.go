//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	pddl "github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD02NullableColumn_ZeroCov(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE TABLE users (\n  id BIGINT PRIMARY KEY,\n  email VARCHAR(255)\n);\n"
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d02NullableColumn(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-2 diag for nullable email column")
	}
	// missing db dir → nil
	if got := d02NullableColumn(&yongol.Fullstack{SpecsDir: t.TempDir()}); got != nil {
		t.Errorf("missing db dir should give nil")
	}
}

func TestD08SerialTypeBanned_ZeroCov(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE TABLE t (\n  id BIGSERIAL PRIMARY KEY\n);\n"
	if err := os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d08SerialTypeBanned(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-8 diag for BIGSERIAL column")
	}
	// missing dir → nil
	if got := d08SerialTypeBanned(&yongol.Fullstack{SpecsDir: t.TempDir()}); got != nil {
		t.Errorf("missing db dir should give nil")
	}
}

func TestD11UnsupportedPgType_ZeroCov(t *testing.T) {
	cols := map[string]pddl.Column{
		"c": {Name: "c", RawType: "TIME WITH TIME ZONE"},
	}
	fs := &yongol.Fullstack{
		DDLTables: []pddl.Table{{Name: "t", File: "t.sql", Line: 1, Columns: cols}},
	}
	diags := d11UnsupportedPgType(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-11 diag for unsupported type")
	}
	if got := d11UnsupportedPgType(nil); got != nil {
		t.Errorf("nil fs should give nil")
	}
}

func TestReadDBSQLFiles_ZeroCov(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "a.sql"), []byte("CREATE TABLE a (id BIGINT);"), 0o644); err != nil {
		t.Fatal(err)
	}
	files := readDBSQLFiles(tmp)
	if len(files) != 1 {
		t.Errorf("expected 1 sql file, got %d", len(files))
	}
	// missing dir → nil
	if got := readDBSQLFiles(filepath.Join(t.TempDir(), "nope")); got != nil {
		t.Errorf("missing dir should give nil")
	}
}

func TestReadSQLDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "x.sql"), []byte("CREATE TABLE x (id BIGINT);"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "skip.txt"), []byte("ignore"), 0o644); err != nil {
		t.Fatal(err)
	}
	// baseline file must be skipped
	if err := os.WriteFile(filepath.Join(dir, ".latest_schema.sql"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := readSQLDir(dir)
	if len(out) != 1 || out[0].name != "x.sql" {
		t.Errorf("expected only x.sql, got %#v", out)
	}
	if got := readSQLDir(filepath.Join(dir, "missing")); got != nil {
		t.Errorf("missing dir should give nil")
	}
}
