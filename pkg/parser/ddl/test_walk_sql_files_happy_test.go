//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what walkSQLFiles — 디렉토리 내 모든 *.sql 파일에 대해 handler 호출

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestWalkSQLFiles_Happy(t *testing.T) {
	dir := t.TempDir()
	files := []string{"a.sql", "b.sql", "c.sql"}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("-- "+name), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var seen []string
	diags := walkSQLFiles(dir, func(path string, data []byte) []diagnostic.Diagnostic {
		seen = append(seen, filepath.Base(path))
		return nil
	})
	if len(diags) != 0 {
		t.Fatalf("diags = %v, want empty", diags)
	}
	if len(seen) != len(files) {
		t.Fatalf("handler called %d times, want %d (seen=%v)", len(seen), len(files), seen)
	}
}
