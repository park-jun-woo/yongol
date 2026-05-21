//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what XFS-01 — 모든 state가 stateDiagram에 있을 때 정상 통과 테스트
package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestXFS01_StatesInDiagram_NoFire(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {States: []string{"draft", "active"}},
	}
	diagrams := []*statemachine.StateDiagram{
		{ID: "workflows", States: []string{"draft", "active", "archived"}},
	}
	fs := buildFSForXFS(ft, diagrams)
	diags := xfs01StatesInDiagram(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
