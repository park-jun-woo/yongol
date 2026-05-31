//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

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
