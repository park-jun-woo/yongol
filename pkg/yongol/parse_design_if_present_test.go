//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseDesignIfPresent — DESIGN.md 미존재(return) + 존재 시 DesignSpec 설정

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDesignIfPresent_Missing(t *testing.T) {
	root := t.TempDir() // no frontend/DESIGN.md
	fs := &Fullstack{Presences: map[SSOTKind]SSOTPresence{}}
	parseDesignIfPresent(fs, root)
	if fs.DesignSpec != nil {
		t.Fatalf("expected no DesignSpec when DESIGN.md missing")
	}
	if _, ok := fs.Presences[KindDesign]; ok {
		t.Fatalf("expected KindDesign not marked present")
	}
}

func TestParseDesignIfPresent_Present(t *testing.T) {
	root := t.TempDir()
	frontend := filepath.Join(root, "frontend")
	if err := os.MkdirAll(frontend, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "---\nversion: \"1.0\"\nname: T\ncomponents:\n  Card:\n    base: \"rounded\"\n---\n"
	if err := os.WriteFile(filepath.Join(frontend, "DESIGN.md"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &Fullstack{Presences: map[SSOTKind]SSOTPresence{}}
	parseDesignIfPresent(fs, root)
	if fs.DesignSpec == nil {
		t.Fatalf("expected DesignSpec to be set")
	}
	if fs.Presences[KindDesign] != SSOTPopulated {
		t.Fatalf("expected KindDesign SSOTPopulated, got %v", fs.Presences[KindDesign])
	}
}
