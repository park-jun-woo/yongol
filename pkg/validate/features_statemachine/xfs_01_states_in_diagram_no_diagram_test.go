//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what XFS-01 — 테이블에 대응하는 stateDiagram이 없을 때 ERROR 테스트
package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestXFS01_StatesInDiagram_NoDiagramForTable(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {States: []string{"draft", "active"}},
	}
	// Diagram exists but for a different entity.
	diagrams := []*statemachine.StateDiagram{
		{ID: "orders", States: []string{"pending", "shipped"}},
	}
	fs := buildFSForXFS(ft, diagrams)
	diags := xfs01StatesInDiagram(fs)
	if len(diags) != 2 {
		t.Fatalf("want 2 diags (one per state), got %d", len(diags))
	}
}
