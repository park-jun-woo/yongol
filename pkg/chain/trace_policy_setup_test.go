//ff:func feature=chain type=test-helper control=sequence
//ff:what tracePolicySetup — authz/project.rego 를 생성하고 (specsDir, regoFile) 반환하는 테스트 헬퍼
package chain

import (
	"os"
	"path/filepath"
	"testing"
)

// tracePolicySetup creates a temp specs dir with an authz/project.rego file and
// returns the specs dir and the rego file path.
func tracePolicySetup(t *testing.T) (specsDir, regoFile string) {
	t.Helper()
	specsDir = t.TempDir()
	authzDir := filepath.Join(specsDir, "authz")
	if err := os.MkdirAll(authzDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	regoFile = filepath.Join(authzDir, "project.rego")
	content := "package project\n\nallow if input.resource == \"project\"\n"
	if err := os.WriteFile(regoFile, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	return specsDir, regoFile
}
