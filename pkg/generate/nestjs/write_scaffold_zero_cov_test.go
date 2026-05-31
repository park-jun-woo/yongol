//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteScaffold_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeScaffold(dir, "myapp"); err != nil {
		t.Fatalf("writeScaffold error: %v", err)
	}
	for _, name := range []string{"package.json", "tsconfig.json", "nest-cli.json"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected %s: %v", name, err)
		}
	}
}
