//ff:func feature=validate type=test control=sequence
//ff:what TestBuildDiagramStateMap_ZeroCov — stateDiagram 배열 → ID별 상태 집합 맵 직접 호출

package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestBuildDiagramStateMap_ZeroCov(t *testing.T) {
	diagrams := []*statemachine.StateDiagram{
		{ID: "order", States: []string{"pending", "shipped"}},
		{ID: "user", States: []string{"active"}},
	}
	m := buildDiagramStateMap(diagrams)
	if len(m) != 2 {
		t.Fatalf("map size = %d", len(m))
	}
	if !m["order"]["pending"] || !m["order"]["shipped"] {
		t.Errorf("order states = %v", m["order"])
	}
	if !m["user"]["active"] {
		t.Errorf("user states = %v", m["user"])
	}
}
