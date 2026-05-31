//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectGuardStateFuncs_ZeroCov(t *testing.T) {
	funcs := []ssac.ServiceFunc{
		{Name: "WithState", Sequences: []ssac.Sequence{{Type: "state", DiagramID: "d"}}},
		{Name: "NoState", Sequences: []ssac.Sequence{{Type: "get"}}},
	}
	m := collectGuardStateFuncs(funcs)
	if !m["WithState"] || m["NoState"] {
		t.Fatalf("collectGuardStateFuncs = %v", m)
	}
}
