//ff:func feature=validate type=test control=sequence topic=tsx
//ff:what T-01 테스트 — 로컬 컴포넌트 import 대상 파일이 존재하는지 (WARNING)

package tsx

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
)

// buildSpecsDir scaffolds a minimal specs/ tree:
//
//	<tmp>/specs/frontend/src/components/ui/Button.tsx
//	<tmp>/specs/frontend/src/pages/Home.tsx
//
// and returns (specsDir, pageFile).
func buildSpecsDir(t *testing.T, existingComponents []string) (string, string) {
	t.Helper()
	dir := t.TempDir()
	specsDir := filepath.Join(dir, "specs")
	pagesDir := filepath.Join(specsDir, "frontend", "src", "pages")
	if err := os.MkdirAll(pagesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, rel := range existingComponents {
		full := filepath.Join(specsDir, "frontend", "src", rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("export const X = 1;"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	pageFile := filepath.Join(pagesDir, "Home.tsx")
	if err := os.WriteFile(pageFile, []byte("export default function Home() { return null }"), 0o644); err != nil {
		t.Fatal(err)
	}
	return specsDir, pageFile
}

func TestT01_ExistingComponent(t *testing.T) {
	specsDir, pageFile := buildSpecsDir(t, []string{"components/ui/Button.tsx"})
	fs := &yongol.Fullstack{
		SpecsDir: specsDir,
		TSXPages: []tsxparser.PageSpec{{
			File:    pageFile,
			Imports: []tsxparser.ComponentImport{{Name: "Button", Path: "@/components/ui/Button", Line: 3}},
		}},
	}
	if diags := t01ComponentFile(fs); len(diags) != 0 {
		t.Fatalf("existing file should produce no diagnostic, got %+v", diags)
	}
}

func TestT01_MissingComponentWarning(t *testing.T) {
	specsDir, pageFile := buildSpecsDir(t, nil)
	fs := &yongol.Fullstack{
		SpecsDir: specsDir,
		TSXPages: []tsxparser.PageSpec{{
			File:    pageFile,
			Imports: []tsxparser.ComponentImport{{Name: "Button", Path: "@/components/ui/Button", Line: 3}},
		}},
	}
	diags := t01ComponentFile(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected WARNING, got %s", diags[0].Level)
	}
	if !strings.Contains(diags[0].Message, "[T-1]") {
		t.Errorf("missing rule id T-1: %q", diags[0].Message)
	}
}

func TestT01_RelativeImport(t *testing.T) {
	specsDir, pageFile := buildSpecsDir(t, []string{"pages/sibling/Widget.tsx"})
	// page imports "./sibling/Widget" from pages/Home.tsx
	fs := &yongol.Fullstack{
		SpecsDir: specsDir,
		TSXPages: []tsxparser.PageSpec{{
			File:    pageFile,
			Imports: []tsxparser.ComponentImport{{Name: "Widget", Path: "./components/sibling/Widget", Line: 2}},
		}},
	}
	// Build the sibling under components/ to match isLocalComponentImport matcher.
	sibling := filepath.Join(specsDir, "frontend", "src", "pages", "components", "sibling", "Widget.tsx")
	if err := os.MkdirAll(filepath.Dir(sibling), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(sibling, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	if diags := t01ComponentFile(fs); len(diags) != 0 {
		t.Fatalf("relative import existing file should produce no diagnostic, got %+v", diags)
	}
}
