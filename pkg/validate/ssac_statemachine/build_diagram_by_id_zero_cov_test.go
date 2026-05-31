//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestBuildDiagramByID_ZeroCov(t *testing.T) {
	m := buildDiagramByID([]*statemachine.StateDiagram{{ID: "a"}, {ID: "b"}})
	if len(m) != 2 || m["a"] == nil || m["b"] == nil {
		t.Fatalf("buildDiagramByID = %v", m)
	}
}
