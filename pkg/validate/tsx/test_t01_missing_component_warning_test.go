//ff:func feature=validate type=test control=sequence topic=tsx
//ff:what T-01 — 존재하지 않는 컴포넌트 import 는 WARNING

package tsx

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
