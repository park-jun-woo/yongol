//ff:func feature=validate type=test control=sequence topic=tsx
//ff:what T-01 — 상대 경로 import 도 실제 파일이 있으면 진단 없음

package tsx

import (
	"os"
	"path/filepath"
	"testing"

	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
