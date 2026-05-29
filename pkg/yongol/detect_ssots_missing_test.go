//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — 빈 specs 루트는 empty slice + no error
package yongol

import (
	"testing"
)

// TestDetectSSOTsEmptyDir asserts that a completely empty specs root returns
// no detected SSOTs and no error. All kinds are SSOTAbsent and therefore
// omitted by design.
func TestDetectSSOTsEmptyDir(t *testing.T) {
	tmp := newTmpSpecsDir(t)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if len(detected) != 0 {
		t.Fatalf("expected 0 detected on empty dir, got %d: %+v", len(detected), detected)
	}
}
