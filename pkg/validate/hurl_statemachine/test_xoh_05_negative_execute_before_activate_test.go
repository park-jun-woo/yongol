//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what TestXoh05_Negative_ExecuteBeforeActivate — 활성화 없이 실행 시 [XOH-05]

package hurl_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh05_Negative_ExecuteBeforeActivate(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc:    newWorkflowDoc(),
		StateDiagrams: []*statemachine.StateDiagram{newWorkflowDiagram()},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
			{Method: "POST", Path: "/workflows/1/execute", File: "t.hurl", Line: 5},
		},
	}
	diags := xoh05StateTransitionOrder(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XOH-05]") || !strings.Contains(diags[0].Message, "ExecuteWorkflow") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}
