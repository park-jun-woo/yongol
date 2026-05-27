//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectOpImportSelfSkip — 동일 feature 패키지 self-import 방지 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectOpImportSelfSkip(t *testing.T) {
	t.Run("SkipSameFeatureCall", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpCall,
			Call: &ir.CallOp{
				Package:  "auth",
				Function: "IssueToken",
			},
		}
		collectOpImport(&d, op, "auth")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("expected no external import for self-reference, got: %v", d.ExtPkgs)
		}
	})

	t.Run("KeepDifferentFeatureCall", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpCall,
			Call: &ir.CallOp{
				Package:  "billing",
				Function: "IsZeroBalance",
			},
		}
		collectOpImport(&d, op, "workflow")
		if _, ok := d.ExtPkgs["billing"]; !ok {
			t.Error("expected billing import for cross-feature reference")
		}
	})

	t.Run("SkipSameFeatureEval", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpEval,
			Eval: &ir.EvalOp{
				Package:  "auth",
				Function: "CheckToken",
			},
		}
		collectOpImport(&d, op, "auth")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("expected no external import for self-reference eval, got: %v", d.ExtPkgs)
		}
	})
}
