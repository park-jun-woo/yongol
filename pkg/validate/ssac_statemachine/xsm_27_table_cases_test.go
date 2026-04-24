//ff:func feature=validate type=test-helper control=sequence topic=states
//ff:what xsm27TableCases — TestXsm27StateIntentDeclaration 테이블 케이스 빌더

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xsm27TableCases returns the full matrix of fixtures exercised by
// TestXsm27StateIntentDeclaration. Extracted so the outer test stays
// within the Q4 PURE line budget.
func xsm27TableCases(findByID, stateSeq, nonStatefulGet ssac.Sequence) []xsm27Case {
	return []xsm27Case{
		{
			name:        "ExecuteWorkflow (POST, stateful, no state, no neutral) fires",
			method:      "POST",
			path:        "/workflows/{id}/execute",
			opID:        "ExecuteWorkflow",
			sequences:   []ssac.Sequence{findByID},
			withDiagram: true,
			withDefault: true,
			wantFire:    true,
		},
		{
			name:        "ActivateWorkflow (POST, with @state) passes",
			method:      "POST",
			path:        "/workflows/{id}/activate",
			opID:        "ActivateWorkflow",
			sequences:   []ssac.Sequence{findByID, stateSeq},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:         "LikeWorkflow (POST, with @state-neutral) passes",
			method:       "POST",
			path:         "/workflows/{id}/like",
			opID:         "LikeWorkflow",
			sequences:    []ssac.Sequence{findByID},
			stateNeutral: true,
			withDiagram:  true,
			withDefault:  true,
			wantFire:     false,
		},
		{
			name:        "ListWorkflows (GET) — method mismatch, passes",
			method:      "GET",
			path:        "/workflows/{id}",
			opID:        "GetWorkflow",
			sequences:   []ssac.Sequence{findByID},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:        "CreateWorkflow (POST, no {id}) — passes",
			method:      "POST",
			path:        "/workflows",
			opID:        "CreateWorkflow",
			sequences:   []ssac.Sequence{},
			withDiagram: true,
			withDefault: true,
			wantFire:    false,
		},
		{
			name:           "DeleteComment (POST on non-stateful resource) — passes",
			method:         "DELETE",
			path:           "/comments/{id}",
			opID:           "DeleteComment",
			sequences:      []ssac.Sequence{nonStatefulGet},
			nonStatefulRes: true,
			wantFire:       false,
		},
	}
}
