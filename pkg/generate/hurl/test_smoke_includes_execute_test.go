//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeIncludesExecute — OpenAPI operationId ExecuteWorkflow 이 smoke 출력에 포함 검증

package hurl

import (
	"testing"
)

// TestSmokeIncludesExecute pins BUG-016 Phase004: the state-dependent
// operation ExecuteWorkflow (active → active self-loop) must show up in
// the generated smoke sequence. Previously the smoke walker collapsed
// per-from-state to a single outgoing transition, silently dropping
// self-loops like Execute and terminal edges like Archive.
func TestSmokeIncludesExecute(t *testing.T) {
	fs := newSmokeFullstack(newZenflowLikeOpenAPI(), newZenflowStateDiagram())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)
	if indexOfString(opIDs, "ExecuteWorkflow") < 0 {
		t.Errorf("ExecuteWorkflow missing from smoke steps: %v", opIDs)
	}
}
