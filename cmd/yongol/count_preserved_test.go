//ff:func feature=cli type=test control=sequence
//ff:what countPreserved test — reason 유무 분류 카운트 검증

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountPreserved(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		w, wo := countPreserved(nil)
		if w != 0 || wo != 0 {
			t.Fatalf("expected (0,0), got (%d,%d)", w, wo)
		}
	})
	t.Run("NonexistentFiles", func(t *testing.T) {
		w, wo := countPreserved([]string{"/tmp/no-such-file-yongol-1", "/tmp/no-such-file-yongol-2"})
		if w != 0 {
			t.Errorf("expected 0 withReason, got %d", w)
		}
		if wo != 2 {
			t.Errorf("expected 2 withoutReason, got %d", wo)
		}
	})
	t.Run("WithReason", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "test.go")
		content := "//ff:preserve reason=\"custom edit\"\npackage foo\n"
		if err := os.WriteFile(f, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		w, wo := countPreserved([]string{f})
		if w != 1 {
			t.Errorf("expected 1 withReason, got %d", w)
		}
		if wo != 0 {
			t.Errorf("expected 0 withoutReason, got %d", wo)
		}
	})
	t.Run("Mixed", func(t *testing.T) {
		dir := t.TempDir()
		withReason := filepath.Join(dir, "with.go")
		withoutReason := filepath.Join(dir, "without.go")
		if err := os.WriteFile(withReason, []byte("//ff:preserve reason=\"edit\"\npackage x\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(withoutReason, []byte("package x\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		w, wo := countPreserved([]string{withReason, withoutReason, "/tmp/nonexist"})
		if w != 1 {
			t.Errorf("expected 1 withReason, got %d", w)
		}
		if wo != 2 {
			t.Errorf("expected 2 withoutReason, got %d", wo)
		}
	})
}
