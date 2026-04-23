//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeStateOrderActivateBeforeExecute — ActivateWorkflow 가 ExecuteWorkflow 보다 먼저 배치 검증

package hurl

import (
	"testing"
)

// TestSmokeStateOrderActivateBeforeExecute pins BUG-016 Phase004:
// ExecuteWorkflow requires the workflow to already be in the `active`
// state, so ActivateWorkflow (draft → active) must precede
// ExecuteWorkflow (active → active) in smoke. The state-machine-aware
// walker guarantees this by BFSing from the initial state.
func TestSmokeStateOrderActivateBeforeExecute(t *testing.T) {
	fs := newSmokeFullstack(newZenflowLikeOpenAPI(), newZenflowStateDiagram())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)

	activateIdx := indexOfString(opIDs, "ActivateWorkflow")
	executeIdx := indexOfString(opIDs, "ExecuteWorkflow")
	if activateIdx < 0 {
		t.Fatalf("ActivateWorkflow missing from smoke steps: %v", opIDs)
	}
	if executeIdx < 0 {
		t.Fatalf("ExecuteWorkflow missing from smoke steps: %v", opIDs)
	}
	if activateIdx > executeIdx {
		t.Errorf("ActivateWorkflow (idx=%d) must precede ExecuteWorkflow (idx=%d); got order=%v",
			activateIdx, executeIdx, opIDs)
	}
}
