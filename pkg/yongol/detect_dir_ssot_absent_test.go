//ff:func feature=orchestrator type=test control=sequence
//ff:what TestDetectDirSSOT/directorySSOTs — glob 매칭 presence 결정 및 후보 목록 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestDetectDirSSOTAbsent(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-dir")
	d := dirSSOT{kind: KindDDL, dir: missing, globs: []string{"*.sql"}}

	got, err := detectDirSSOT(d)
	if err != nil {
		t.Fatal(err)
	}
	if got.Presence != SSOTAbsent {
		t.Errorf("Presence = %v, want SSOTAbsent", got.Presence)
	}
	// Absent result carries no Kind/Path.
	if got.Kind != "" {
		t.Errorf("expected empty Kind for absent, got %q", got.Kind)
	}
}
