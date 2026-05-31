//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildFuncByName_ZeroCov(t *testing.T) {
	m := buildFuncByName([]ssac.ServiceFunc{{Name: "X"}, {Name: "Y"}})
	if len(m) != 2 || m["X"].Name != "X" {
		t.Fatalf("buildFuncByName = %v", m)
	}
}
