//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — db/*.sql populated 분기
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsDDLPopulated covers the db/*.sql happy path: directory exists
// and contains at least one SQL file → SSOTPopulated.
func TestDetectSSOTsDDLPopulated(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "db", "users.sql"), "CREATE TABLE users (id INT);\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindDDL)
	if !ok {
		t.Fatalf("KindDDL not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated, got %v", d.Presence)
	}
}
