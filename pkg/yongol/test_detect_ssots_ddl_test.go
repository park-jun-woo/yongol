//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — db/ *.sql presence / declared-empty 분기
package yongol

import (
	"os"
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
