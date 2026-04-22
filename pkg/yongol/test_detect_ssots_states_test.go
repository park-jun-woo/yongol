//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — states/*.md presence 감지
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsStatesPopulated verifies that a Mermaid state diagram under
// states/ marks KindStates as SSOTPopulated.
func TestDetectSSOTsStatesPopulated(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "states", "Workflow.md"),
		"```mermaid\nstateDiagram-v2\n[*] --> Draft\n```\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindStates)
	if !ok {
		t.Fatalf("KindStates not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated, got %v", d.Presence)
	}
}
