//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeOrderingCreateBeforeMutate — POST /workflows가 POST /workflows/{id}/activate 보다 먼저 배치 검증

package hurl

import (
	"testing"
)

// TestSmokeOrderingCreateBeforeMutate pins Phase003: for a SSOT with
// `POST /workflows` (create) + `POST /workflows/{id}/activate` (mutate),
// the creation step must appear before the mutation. Otherwise the
// mutation has no captured `{id}` to substitute into the path.
func TestSmokeOrderingCreateBeforeMutate(t *testing.T) {
	fs := newSmokeFullstack(newZenflowLikeOpenAPI(), newZenflowStateDiagram())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)

	createIdx := indexOfString(opIDs, "CreateWorkflow")
	activateIdx := indexOfString(opIDs, "ActivateWorkflow")
	if createIdx < 0 {
		t.Fatalf("CreateWorkflow missing from smoke steps: %v", opIDs)
	}
	if activateIdx < 0 {
		t.Fatalf("ActivateWorkflow missing from smoke steps: %v", opIDs)
	}
	if createIdx > activateIdx {
		t.Errorf("CreateWorkflow (idx=%d) must precede ActivateWorkflow (idx=%d); got order=%v",
			createIdx, activateIdx, opIDs)
	}
}
