//ff:func feature=ddl type=test control=sequence
//ff:what walkSQLFiles — *.sql 외 확장자와 하위 디렉토리는 제외

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestWalkSQLFiles_MixedFilter(t *testing.T) {
	dir := t.TempDir()
	// .sql → included
	if err := os.WriteFile(filepath.Join(dir, "users.sql"), []byte("-- sql"), 0o644); err != nil {
		t.Fatal(err)
	}
	// .txt → excluded
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("text"), 0o644); err != nil {
		t.Fatal(err)
	}
	// subdirectory → excluded (non-recursive)
	sub := filepath.Join(dir, "sub")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "nested.sql"), []byte("-- nested"), 0o644); err != nil {
		t.Fatal(err)
	}
	var seen []string
	diags := walkSQLFiles(dir, func(path string, data []byte) []diagnostic.Diagnostic {
		seen = append(seen, filepath.Base(path))
		return nil
	})
	if len(diags) != 0 {
		t.Fatalf("diags = %v, want empty", diags)
	}
	if len(seen) != 1 || seen[0] != "users.sql" {
		t.Fatalf("seen = %v, want [users.sql]", seen)
	}
}
