//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildFuncNameSet_ZeroCov(t *testing.T) {
	m := buildFuncNameSet([]ssac.ServiceFunc{{Name: "X"}})
	if !m["X"] || m["Z"] {
		t.Fatalf("buildFuncNameSet = %v", m)
	}
}
