//ff:func feature=manifest type=test control=sequence
//ff:what ParseDir — 정상 DDL 파일은 1개 결과, 0개 진단 반환

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDir_Happy(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "users.sql")
	if err := os.WriteFile(path, []byte("CREATE TABLE users (id BIGSERIAL PRIMARY KEY);"), 0o644); err != nil {
		t.Fatal(err)
	}
	results, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(results) != 1 {
		t.Fatalf("results count = %d, want 1", len(results))
	}
}
