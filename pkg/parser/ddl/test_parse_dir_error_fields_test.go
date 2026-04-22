//ff:func feature=manifest type=test control=sequence
//ff:what ParseDir — 존재하지 않는 디렉토리에서 Diagnostic 필드 완결성 회귀

package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseDir_MissingDir_DiagnosticFields(t *testing.T) {
	_, diags := ParseDir("/nonexistent/ddl/xyz123")
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.File != "/nonexistent/ddl/xyz123" {
		t.Errorf("File = %q, want dir", d.File)
	}
}

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
