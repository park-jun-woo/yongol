//ff:func feature=cli type=test control=sequence
//ff:what printPreservedList test — Preserved Files 섹션 출력 검증

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrintPreservedList(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var buf bytes.Buffer
		printPreservedList(&buf, nil)
		out := buf.String()
		if !strings.Contains(out, "Preserved Files (0)") {
			t.Errorf("expected 0 count, got: %q", out)
		}
	})

	t.Run("WithReason", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "svc.go")
		if err := os.WriteFile(f, []byte("//ff:preserve reason=\"user hook\"\npackage svc\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		printPreservedList(&buf, []string{f})
		out := buf.String()

		if !strings.Contains(out, "Preserved Files (1)") {
			t.Errorf("expected 1 file, got: %q", out)
		}
		if !strings.Contains(out, "user hook") {
			t.Errorf("expected reason, got: %q", out)
		}
	})

	t.Run("WithoutReason", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "plain.go")
		if err := os.WriteFile(f, []byte("package plain\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		printPreservedList(&buf, []string{f})
		out := buf.String()

		if !strings.Contains(out, "<none>") {
			t.Errorf("expected <none> reason, got: %q", out)
		}
	})

	t.Run("NonexistentFile", func(t *testing.T) {
		var buf bytes.Buffer
		printPreservedList(&buf, []string{"/tmp/no-such-yongol-file.go"})
		out := buf.String()

		if !strings.Contains(out, "<none>") {
			t.Errorf("expected <none> for missing file, got: %q", out)
		}
	})
}
