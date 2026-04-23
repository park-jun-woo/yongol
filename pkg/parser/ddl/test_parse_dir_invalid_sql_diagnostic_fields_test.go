//ff:func feature=manifest type=test control=sequence
//ff:what ParseDir — 잘못된 SQL 구문은 File/Phase/Level 진단 필드와 함께 보고

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseDir_InvalidSQL_DiagnosticFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.sql")
	if err := os.WriteFile(path, []byte("CREATE TABLE ((( invalid"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseDir(dir)
	if len(diags) < 1 {
		t.Fatalf("expected parse error diag, got none")
	}
	d := diags[0]
	if d.File != path {
		t.Errorf("File = %q, want %q", d.File, path)
	}
	if d.Phase != diagnostic.PhaseParse || d.Level != diagnostic.LevelError {
		t.Errorf("phase/level missing: %+v", d)
	}
}
