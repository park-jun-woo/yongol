//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what writeComponentsUI fallback 프리미티브 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteComponentsUI_FallbackPrimitives(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := writeComponentsUI(srcDir, nil); err != nil {
		t.Fatal(err)
	}
	uiDir := filepath.Join(srcDir, "components", "ui")
	for name := range uiPrimitives() {
		path := filepath.Join(uiDir, name)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected fallback primitive file %s", name)
		}
	}
}
