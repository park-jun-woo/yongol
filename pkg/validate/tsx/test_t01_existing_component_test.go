//ff:func feature=validate type=test control=sequence topic=tsx
//ff:what T-01 — 실제 존재하는 컴포넌트 import 는 진단 없음

package tsx

import (
	"testing"

	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
