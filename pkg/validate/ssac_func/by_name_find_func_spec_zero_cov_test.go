//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestByName_ZeroCov — ssac_func 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

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
