//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestXsm27StateIntentDeclaration — @state / @state-neutral 강제 테이블 기반 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestXsm27StateIntentDeclaration exercises XSM-27 over a matrix of
// (method, path, @state / @state-neutral / plain) combinations to pin the
// gate conditions documented on xsm27StateIntentDeclaration.
func TestXsm27StateIntentDeclaration(t *testing.T) {
	findByID := ssac.Sequence{
		Type:  "get",
		Model: "Workflow.FindByID",
		Args: []ssac.Arg{
			{Source: "request", Field: "id"},
		},
		Result: &ssac.Result{Type: "Workflow", Var: "wf"},
	}
	stateSeq := ssac.Sequence{
		Type:       "state",
		DiagramID:  "workflow",
		Inputs:     map[string]string{"Status": "wf.Status"},
		Transition: "ActivateWorkflow",
		Message:    "Cannot activate",
	}
	nonStatefulGet := ssac.Sequence{
		Type:  "get",
		Model: "Comment.FindByID",
		Args: []ssac.Arg{
			{Source: "request", Field: "id"},
		},
		Result: &ssac.Result{Type: "Comment", Var: "c"},
	}

	cases := xsm27TableCases(findByID, stateSeq, nonStatefulGet)

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			runXsm27Case(t, tc)
		})
	}
}
