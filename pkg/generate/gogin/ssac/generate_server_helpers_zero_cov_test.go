//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestGenerateServerHelpers_ZeroCov — ptr_of/deref_* 헬퍼 파일 emit
package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateServerHelpers_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := generateServerHelpers(dir); err != nil {
		t.Fatalf("generateServerHelpers: %v", err)
	}
	serviceDir := filepath.Join(dir, "backend", "internal", "service")
	entries, err := os.ReadDir(serviceDir)
	if err != nil {
		t.Fatalf("read service dir: %v", err)
	}
	if len(entries) == 0 {
		t.Fatalf("expected helper files emitted")
	}
	foundPtrOf := false
	for _, e := range entries {
		if e.Name() == "ptr_of.go" {
			foundPtrOf = true
		}
	}
	if !foundPtrOf {
		t.Errorf("expected ptr_of.go, got %v", entries)
	}
}
