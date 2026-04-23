//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestSmokeStateOrderArchiveLast — ArchiveWorkflow 가 workflow state-transition 중 마지막 검증

package hurl

import (
	"testing"
)

// TestSmokeStateOrderArchiveLast pins BUG-016 Phase004: Archive is a
// terminal transition (active → archived) that locks the workflow out
// of every other state operation, so among the workflow state events
// Archive must be the last one emitted. Verified by walking through the
// set of state-machine operationIds and checking Archive has the
// greatest index.
func TestSmokeStateOrderArchiveLast(t *testing.T) {
	fs := newSmokeFullstack(newZenflowLikeOpenAPI(), newZenflowStateDiagram())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)

	stateOps := []string{"ActivateWorkflow", "PauseWorkflow", "ExecuteWorkflow", "ArchiveWorkflow"}
	archiveIdx := indexOfString(opIDs, "ArchiveWorkflow")
	if archiveIdx < 0 {
		t.Fatalf("ArchiveWorkflow missing from smoke steps: %v", opIDs)
	}
	for _, op := range stateOps {
		if op == "ArchiveWorkflow" {
			continue
		}
		idx := indexOfString(opIDs, op)
		if idx < 0 {
			t.Errorf("%s missing from smoke steps: %v", op, opIDs)
			continue
		}
		if idx > archiveIdx {
			t.Errorf("%s (idx=%d) must precede ArchiveWorkflow (idx=%d); got order=%v",
				op, idx, archiveIdx, opIDs)
		}
	}
}
