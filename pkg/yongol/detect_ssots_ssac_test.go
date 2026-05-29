//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — service/*.ssac + service/<mod>/*.ssac 두 가지 glob 경로 회귀
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsSSaCNested asserts that the nested glob (service/<mod>/*.ssac)
// still marks the directory SSOTPopulated. zenflow only uses the nested form,
// so this specifically guards that a top-level *.ssac is not required.
func TestDetectSSOTsSSaCNested(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "service", "user", "login.ssac"),
		"@service user\n@method Login\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindSSaC)
	if !ok {
		t.Fatalf("KindSSaC not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated, got %v", d.Presence)
	}
}
