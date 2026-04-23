//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — db/ 디렉토리만 존재(내용 없음) 시 SSOTDeclared
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDetectSSOTsDDLDeclared covers the directory-but-no-content case:
// `db/` exists but no *.sql inside → SSOTDeclared (user signaled intent but
// never populated). This separates "opt-out" from "WIP".
func TestDetectSSOTsDDLDeclared(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	if err := os.MkdirAll(filepath.Join(tmp, "db"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindDDL)
	if !ok {
		t.Fatalf("KindDDL not detected (expected declared); detected=%+v", detected)
	}
	if d.Presence != SSOTDeclared {
		t.Fatalf("expected SSOTDeclared, got %v", d.Presence)
	}
}
