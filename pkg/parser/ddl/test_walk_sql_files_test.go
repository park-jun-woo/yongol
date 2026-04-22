//ff:func feature=ddl type=test control=sequence
//ff:what walkSQLFiles — happy / missing-dir / mixed-filter 3 case regression

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

func TestWalkSQLFiles_MissingDir(t *testing.T) {
	called := false
	diags := walkSQLFiles("/nonexistent/ddl/walksql/xyz123", func(path string, data []byte) []diagnostic.Diagnostic {
		called = true
		return nil
	})
	if called {
		t.Fatalf("handler must not be called when dir read fails")
	}
	if len(diags) != 1 {
		t.Fatalf("diags count = %d, want 1", len(diags))
	}
	d := diags[0]
	if d.File != "/nonexistent/ddl/walksql/xyz123" {
		t.Errorf("File = %q, want dir path", d.File)
	}
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want LevelError", d.Level)
	}
	if d.Message == "" {
		t.Errorf("Message empty")
	}
}

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
