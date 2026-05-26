//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what assertEvalSpec — lookupEvalSpec 테스트 assertion 헬퍼

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func assertEvalSpec(t *testing.T, model string, got *funcspec.FuncSpec, wantNil bool, wantPkg, wantFn string) {
	t.Helper()
	if wantNil {
		if got != nil {
			t.Errorf("lookupEvalSpec(%q) = %+v, want nil", model, got)
		}
		return
	}
	if got == nil {
		t.Fatalf("lookupEvalSpec(%q) = nil, want non-nil", model)
	}
	if got.Package != wantPkg || got.Name != wantFn {
		t.Errorf("lookupEvalSpec(%q) = {%q,%q}, want {%q,%q}", model, got.Package, got.Name, wantPkg, wantFn)
	}
}
