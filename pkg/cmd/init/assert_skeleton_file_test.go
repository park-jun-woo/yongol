//ff:func feature=cli-init type=test-helper control=sequence
//ff:what assertSkeletonFile — stat one file and assert non-empty size

package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func assertSkeletonFile(t *testing.T, target, rel string) {
	t.Helper()
	path := filepath.Join(target, rel)
	info, err := os.Stat(path)
	if err != nil {
		t.Errorf("expected file %s missing: %v", rel, err)
		return
	}
	if info.Size() == 0 {
		t.Errorf("expected file %s is empty", rel)
	}
}
