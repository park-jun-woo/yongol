//ff:func feature=cli-init type=test-helper control=iteration dimension=1
//ff:what assertSkeletonDirs — verifies each expected skeleton directory exists

package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func assertSkeletonDirs(t *testing.T, target string) {
	t.Helper()
	expectDirs := []string{
		"specs/db/queries",
		"specs/service",
		"specs/states",
		"specs/frontend/pages",
		"specs/frontend/components",
		"specs/tests",
	}
	for _, rel := range expectDirs {
		info, err := os.Stat(filepath.Join(target, rel))
		if err != nil || !info.IsDir() {
			t.Errorf("expected directory %s missing", rel)
		}
	}
}
