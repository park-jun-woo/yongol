//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateDomainValidators_WriteError — .go 타겟이 디렉토리일 때 WriteFile 에러 표면화 (BUG-142)

package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateDomainValidators_WriteError(t *testing.T) {
	// Spec copy succeeds but the .go target is a directory → WriteFile fails.
	fs := domainedValidatorFS(t)
	mwDir := t.TempDir()
	// Domains iterate sorted: "admin" first.
	if err := os.Mkdir(filepath.Join(mwDir, "request_validator_admin.go"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := generateDomainValidators(fs, mwDir); err == nil {
		t.Fatal("expected write error, got nil")
	}
}
