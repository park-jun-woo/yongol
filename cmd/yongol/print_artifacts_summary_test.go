//ff:func feature=cli type=test control=sequence
//ff:what printArtifactsSummary test — preserved/drift/generated 카운트 출력 검증

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrintArtifactsSummary(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		arts := t.TempDir()
		backend := filepath.Join(arts, "backend")
		if err := os.MkdirAll(backend, 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(backend, "main.go"))
		mustWriteEmpty(t, filepath.Join(backend, "svc.go"))
		mustWriteEmpty(t, filepath.Join(backend, "handler.go"))

		var buf bytes.Buffer
		printArtifactsSummary(&buf, arts, 1, 0)
		out := buf.String()

		if !strings.Contains(out, "files        3") {
			t.Errorf("expected total=3, got: %s", out)
		}
		if !strings.Contains(out, "generated=2") {
			t.Errorf("expected generated=2, got: %s", out)
		}
		if !strings.Contains(out, "preserved=1") {
			t.Errorf("expected preserved=1, got: %s", out)
		}
	})

	t.Run("PreservedExceedsTotalClampedToZero", func(t *testing.T) {
		arts := t.TempDir()
		backend := filepath.Join(arts, "backend")
		if err := os.MkdirAll(backend, 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(backend, "main.go"))

		var buf bytes.Buffer
		printArtifactsSummary(&buf, arts, 5, 2)
		out := buf.String()

		if !strings.Contains(out, "generated=0") {
			t.Errorf("expected generated clamped to 0, got: %s", out)
		}
		if !strings.Contains(out, "drift=2") {
			t.Errorf("expected drift=2, got: %s", out)
		}
	})

	t.Run("EmptyDir", func(t *testing.T) {
		arts := t.TempDir()

		var buf bytes.Buffer
		printArtifactsSummary(&buf, arts, 0, 0)
		out := buf.String()

		if !strings.Contains(out, "files        0") {
			t.Errorf("expected total=0, got: %s", out)
		}
		if !strings.Contains(out, "frontend     -") {
			t.Errorf("expected frontend placeholder, got: %s", out)
		}
	})
}
