//ff:func feature=orchestrator type=test control=sequence
//ff:what TestDetectDirSSOT/directorySSOTs — glob 매칭 presence 결정 및 후보 목록 검증
package yongol

import (
	"testing"
)

func TestDetectDirSSOTDeclared(t *testing.T) {
	dir := t.TempDir() // exists but no *.sql
	d := dirSSOT{kind: KindDDL, dir: dir, globs: []string{"*.sql"}}

	got, err := detectDirSSOT(d)
	if err != nil {
		t.Fatal(err)
	}
	if got.Presence != SSOTDeclared {
		t.Errorf("Presence = %v, want SSOTDeclared", got.Presence)
	}
	if got.Kind != KindDDL {
		t.Errorf("Kind = %v, want KindDDL", got.Kind)
	}
}
