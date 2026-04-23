//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeIncludesArchive — OpenAPI operationId ArchiveWorkflow 이 smoke 출력에 포함 검증

package hurl

import (
	"testing"
)

// TestSmokeIncludesArchive pins BUG-016 Phase004: the terminal workflow
// transition ArchiveWorkflow (active → archived) must be emitted by the
// smoke walker. Archive used to be dropped because the walker would
// only emit one outgoing transition per from-state.
func TestSmokeIncludesArchive(t *testing.T) {
	fs := newSmokeFullstack(newZenflowLikeOpenAPI(), newZenflowStateDiagram())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)
	if indexOfString(opIDs, "ArchiveWorkflow") < 0 {
		t.Errorf("ArchiveWorkflow missing from smoke steps: %v", opIDs)
	}
}
