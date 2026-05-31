//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestByName_ZeroCov — ssac_func 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
