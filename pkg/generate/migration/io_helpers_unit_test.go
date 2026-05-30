//ff:func feature=migration type=test control=iteration dimension=1
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func sampleSchema() *Schema {
	s := NewSchema()
	t := ensureTable(s, "users")
	t.Columns = []*Column{
		{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
		{Name: "email", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
	}
	t.PrimaryKey = []string{"id"}
	return s
}

func TestWriteAndLoadSnapshot_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", ".latest_schema.sql")
	s := sampleSchema()

	if err := WriteSnapshot(path, s, "v1.2.3", time.Unix(0, 0)); err != nil {
		t.Fatalf("WriteSnapshot error: %v", err)
	}
	// file written + header present
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read written snapshot: %v", err)
	}
	if !strings.HasPrefix(string(data), SnapshotHashHeaderPrefix) {
		t.Errorf("snapshot missing hash header: %q", string(data)[:40])
	}

	loaded, err := LoadSnapshot(path)
	if err != nil {
		t.Fatalf("LoadSnapshot error: %v", err)
	}
	if loaded == nil || loaded.Tables["users"] == nil {
		t.Fatalf("loaded schema missing users table: %+v", loaded)
	}
	if len(loaded.Tables["users"].Columns) != 2 {
		t.Errorf("round-trip lost columns: %+v", loaded.Tables["users"].Columns)
	}
}

func TestLoadSnapshot_Missing(t *testing.T) {
	s, err := LoadSnapshot(filepath.Join(t.TempDir(), "nope.sql"))
	if err != nil {
		t.Errorf("missing file should yield (nil, nil), got err %v", err)
	}
	if s != nil {
		t.Errorf("missing file should yield nil schema, got %+v", s)
	}
}

func TestLoadSnapshot_Drift(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".latest_schema.sql")
	if err := WriteSnapshot(path, sampleSchema(), "v1", time.Unix(0, 0)); err != nil {
		t.Fatalf("WriteSnapshot: %v", err)
	}
	// tamper with the body so the stored hash no longer matches
	data, _ := os.ReadFile(path)
	tampered := string(data) + "\n-- injected drift\n"
	if err := os.WriteFile(path, []byte(tampered), 0644); err != nil {
		t.Fatalf("rewrite: %v", err)
	}
	if _, err := LoadSnapshot(path); err == nil || !strings.Contains(err.Error(), "drift") {
		t.Errorf("expected drift hash-mismatch error, got %v", err)
	}
}

func TestLoadSnapshot_BadHeader(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.sql")
	if err := os.WriteFile(path, []byte("no header prefix here\nbody\n"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := LoadSnapshot(path); err == nil {
		t.Errorf("expected error for missing hash header prefix")
	}

	// single line, no newline -> "no header line"
	path2 := filepath.Join(dir, "oneline.sql")
	if err := os.WriteFile(path2, []byte("oneline"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := LoadSnapshot(path2); err == nil {
		t.Errorf("expected error for snapshot with no header line")
	}
}

func TestNextSequenceNumber(t *testing.T) {
	// missing dir -> 1
	if n, err := NextSequenceNumber(filepath.Join(t.TempDir(), "nope")); err != nil || n != 1 {
		t.Errorf("missing dir -> (%d,%v), want (1,nil)", n, err)
	}

	dir := t.TempDir()
	// empty dir -> 1
	if n, err := NextSequenceNumber(dir); err != nil || n != 1 {
		t.Errorf("empty dir -> (%d,%v), want (1,nil)", n, err)
	}

	// with files: max 0003 .up.sql, plus noise -> 4
	for _, name := range []string{
		"0001_initial.up.sql",
		"0001_initial.down.sql", // down ignored
		"0003_add.up.sql",
		"readme.txt",      // non-sql ignored
		"noprefix.up.sql", // no leading number index 0 -> i<=0 skip
		"02x_bad.up.sql",  // non-numeric prefix -> skip
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	n, err := NextSequenceNumber(dir)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 4 {
		t.Errorf("NextSequenceNumber = %d, want 4", n)
	}
}

func TestLoadDataMigrationSQL(t *testing.T) {
	// nil hints / empty -> nil, nil
	if out, missing := LoadDataMigrationSQL("/tmp", nil); out != nil || missing != nil {
		t.Errorf("nil hints -> (%v,%v), want (nil,nil)", out, missing)
	}

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "mig_users.sql"), []byte("UPDATE users SET x=1;"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	hints := &Hints{DataMigrations: map[string]string{
		"users":  "mig_users.sql", // relative -> resolved against specsDir
		"orders": "missing.sql",   // missing -> reported
	}}
	out, missing := LoadDataMigrationSQL(dir, hints)
	if out["users"] != "UPDATE users SET x=1;" {
		t.Errorf("users sidecar not loaded: %q", out["users"])
	}
	if len(missing) != 1 || missing[0] != "missing.sql" {
		t.Errorf("missing = %v, want [missing.sql]", missing)
	}
}

func TestListSQLFiles(t *testing.T) {
	// missing dir -> (nil, nil)
	if files, err := listSQLFiles(filepath.Join(t.TempDir(), "nope"), nil); err != nil || files != nil {
		t.Errorf("missing dir -> (%v,%v), want (nil,nil)", files, err)
	}

	dir := t.TempDir()
	for _, name := range []string{"b.sql", "a.sql", "skip.sql", "notsql.txt"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	files, err := listSQLFiles(dir, []string{"skip.sql"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("got %d files, want 2: %v", len(files), files)
	}
	// sorted: a.sql before b.sql
	if filepath.Base(files[0]) != "a.sql" || filepath.Base(files[1]) != "b.sql" {
		t.Errorf("files not sorted/filtered correctly: %v", files)
	}
}

func TestLoadPrevSnapshot_Initial(t *testing.T) {
	dir := t.TempDir()
	prev, mode, diags := loadPrevSnapshot(filepath.Join(dir, "nope.sql"), dir)
	if mode != ModeInitial {
		t.Errorf("missing snapshot -> mode %q, want initial", mode)
	}
	if prev == nil || prev.Tables == nil {
		t.Errorf("initial prev should be empty non-nil schema")
	}
	if len(diags) != 0 {
		t.Errorf("clean initial should produce no diags, got %v", diags)
	}
}

func TestLoadPrevSnapshot_StateInconsistent(t *testing.T) {
	dir := t.TempDir()
	migDir := filepath.Join(dir, MigrationsSubdir)
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migDir, "0001_x.up.sql"), []byte("x"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	_, mode, diags := loadPrevSnapshot(filepath.Join(dir, "absent.sql"), dir)
	if mode != ModeInitial {
		t.Errorf("mode = %q, want initial", mode)
	}
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "MIG-006") {
		t.Errorf("expected one MIG-006 diag, got %v", diags)
	}
}

func TestLoadPrevSnapshot_LoadError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.sql")
	// valid header prefix but wrong hash -> LoadSnapshot returns error
	if err := os.WriteFile(path, []byte(SnapshotHashHeaderPrefix+"deadbeef\nbody\n"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	_, _, diags := loadPrevSnapshot(path, dir)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "MIG-006") {
		t.Errorf("expected one MIG-006 load-failed diag, got %v", diags)
	}
}
