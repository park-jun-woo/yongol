//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

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
