//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what XFS-01 — features table의 state가 stateDiagram에 없을 때 ERROR 진단 테스트
package features_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestXFS01_StatesInDiagram_Fires(t *testing.T) {
	ft := map[string]features.TableDef{
		"workflows": {States: []string{"draft", "active", "archived"}},
	}
	diagrams := []*statemachine.StateDiagram{
		{ID: "workflows", States: []string{"draft", "active"}},
	}
	fs := buildFSForXFS(ft, diagrams)
	diags := xfs01StatesInDiagram(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XFS-01]") {
		t.Errorf("want [XFS-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "archived") {
		t.Errorf("want state name in message, got %s", diags[0].Message)
	}
}
