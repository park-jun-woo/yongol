//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestEmbedInSyncWithRootRulebook — 내장 rulebook.md 사본이 루트 rulebook.md 와 byte-equal 인지 가드
package catalog

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// TestEmbedInSyncWithRootRulebook guards against drift between the repo root
// rulebook.md (the canonical document) and the //go:embed copy in this
// package. go:embed cannot reference parent paths (and refuses symlinks as
// irregular files), so the copy must be refreshed via go:generate whenever
// the root rulebook changes. A read failure is a hard failure, not a skip,
// so the guard cannot be silently disabled.
func TestEmbedInSyncWithRootRulebook(t *testing.T) {
	root, err := os.ReadFile(filepath.Join("..", "..", "..", "rulebook.md"))
	if err != nil {
		t.Fatalf("read repo root rulebook.md: %v", err)
	}
	if !bytes.Equal(root, Source()) {
		t.Fatalf("pkg/rule/catalog/rulebook.md is out of sync with repo root rulebook.md;\n" +
			"run: go generate ./pkg/rule/catalog")
	}
}
