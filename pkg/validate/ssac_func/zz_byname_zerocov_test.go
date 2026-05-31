//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestByName_ZeroCov — ssac_func 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameCheckAuthInputType_ZeroCov(t *testing.T) {
	// Empty Inputs → no per-input resolution; function is exercised by name.
	fn := parsessac.ServiceFunc{Name: "GetItem", FileName: "f.ssac"}
	seq := parsessac.Sequence{Line: 5, Inputs: map[string]string{}}
	if d := checkAuthInputType(nil, fn, seq); len(d) != 0 {
		t.Errorf("checkAuthInputType empty inputs should yield no diagnostics, got %v", d)
	}
}

func TestByNameFindFuncSpec_ZeroCov(t *testing.T) {
	project := []funcspec.FuncSpec{{Package: "svc", Name: "Create"}}
	yongolPkg := []funcspec.FuncSpec{{Package: "auth", Name: "Login"}}

	if got := findFuncSpec("svc.Create", project, yongolPkg); got == nil {
		t.Errorf("findFuncSpec project miss")
	}
	if got := findFuncSpec("auth.Login", project, yongolPkg); got == nil {
		t.Errorf("findFuncSpec yongolPkg miss")
	}
	if got := findFuncSpec("nope.X", project, yongolPkg); got != nil {
		t.Errorf("findFuncSpec unexpected hit: %+v", got)
	}
}

func TestByNameMakeAuthTypeDiag_ZeroCov(t *testing.T) {
	// string-compatible source → nil.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "string", "Op"); d != nil {
		t.Errorf("makeAuthTypeDiag string-compatible should be nil")
	}
	// empty source type → nil.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "", "Op"); d != nil {
		t.Errorf("makeAuthTypeDiag empty should be nil")
	}
	// incompatible (uuid) → diagnostic.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "uuid.UUID", "Op"); d == nil {
		t.Errorf("makeAuthTypeDiag incompatible should produce diagnostic")
	}
}

func TestByNameXFS39Advice_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{}

	// no dot → generic advice.
	if got := xfs39Advice("plainfunc", fs); got == "" {
		t.Errorf("xfs39Advice no-dot empty")
	}
	// non-builtin package → generic advice.
	if got := xfs39Advice("custom.Do", fs); got == "" {
		t.Errorf("xfs39Advice non-builtin empty")
	}
	// builtin package with no loaded funcs → exercises collectBuiltinFuncNames /
	// collectGroundFuncNames (nil Ground short-circuits).
	if got := xfs39Advice("auth.Login", fs); got == "" {
		t.Errorf("xfs39Advice builtin empty")
	}

	// direct call to collectBuiltinFuncNames for an empty package.
	_ = collectBuiltinFuncNames("auth", fs)
	seen := map[string]bool{}
	collectGroundFuncNames("auth", fs, seen)
}
