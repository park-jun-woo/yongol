//ff:func feature=chain type=test control=sequence
//ff:what findDDLTable 가 db/*.sql 에서 CREATE TABLE 위치를 찾고 미발견/디렉토리없음을 처리하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindDDLTable(t *testing.T) {
	specsDir := t.TempDir()

	// No db/ directory → fallback.
	if rel, line := findDDLTable("courses", specsDir); rel != "db/?.sql" || line != 0 {
		t.Errorf("no db dir: got (%q, %d), want (db/?.sql, 0)", rel, line)
	}

	dbDir := filepath.Join(specsDir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	sql := "-- header\nCREATE TABLE courses (\n  id BIGINT\n);\n"
	if err := os.WriteFile(filepath.Join(dbDir, "schema.sql"), []byte(sql), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	// A non-sql file must be ignored.
	if err := os.WriteFile(filepath.Join(dbDir, "notes.txt"), []byte("CREATE TABLE courses"), 0o644); err != nil {
		t.Fatalf("write txt: %v", err)
	}

	if rel, line := findDDLTable("courses", specsDir); rel != "db/schema.sql" || line != 2 {
		t.Errorf("found table: got (%q, %d), want (db/schema.sql, 2)", rel, line)
	}

	// Table not present → fallback.
	if rel, line := findDDLTable("missing", specsDir); rel != "db/?.sql" || line != 0 {
		t.Errorf("missing table: got (%q, %d), want (db/?.sql, 0)", rel, line)
	}

	// A *.sql entry that cannot be opened (a dangling symlink) exercises the
	// os.Open error branch (continue). The scan still finds the table in the
	// valid schema.sql afterwards.
	t.Run("UnreadableSQLEntry", func(t *testing.T) {
		dir := t.TempDir()
		db := filepath.Join(dir, "db")
		if err := os.MkdirAll(db, 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
		// Dangling symlink named *.sql: ReadDir lists it, IsDir()==false,
		// HasSuffix(".sql")==true, but os.Open fails → continue.
		if err := os.Symlink(filepath.Join(db, "nonexistent-target.sql"), filepath.Join(db, "broken.sql")); err != nil {
			t.Skipf("symlink unsupported: %v", err)
		}
		// Use a "zname" so the valid file sorts after the broken symlink and the
		// continue branch is taken before the match is found.
		good := "-- h\nCREATE TABLE courses (\n  id BIGINT\n);\n"
		if err := os.WriteFile(filepath.Join(db, "zschema.sql"), []byte(good), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
		rel, line := findDDLTable("courses", dir)
		if rel != "db/zschema.sql" || line != 2 {
			t.Errorf("with broken symlink: got (%q, %d), want (db/zschema.sql, 2)", rel, line)
		}
	})
}
