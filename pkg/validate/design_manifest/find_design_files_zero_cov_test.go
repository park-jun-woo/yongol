//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestByName_ZeroCov — normPath / findDesignFiles 직접 호출
package design_manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindDesignFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"DESIGN.md", "theme.design.md", "other.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	got := findDesignFiles(dir)
	if len(got) != 2 {
		t.Errorf("findDesignFiles = %v, want 2 (DESIGN.md + *.design.md)", got)
	}

	// empty dir → no matches.
	if got := findDesignFiles(t.TempDir()); len(got) != 0 {
		t.Errorf("empty dir should yield no matches, got %v", got)
	}
}
