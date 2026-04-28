//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what TestXoh05_Positive_ActivateThenExecute — 활성화 후 실행 시 진단 없음

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh05_Positive_ActivateThenExecute(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc:    newWorkflowDoc(),
		StateDiagrams: []*statemachine.StateDiagram{newWorkflowDiagram()},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
			{Method: "POST", Path: "/workflows/1/activate", File: "t.hurl", Line: 3},
			{Method: "POST", Path: "/workflows/1/execute", File: "t.hurl", Line: 5},
		},
	}
	if diags := xoh05StateTransitionOrder(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
