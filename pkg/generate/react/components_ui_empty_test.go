//ff:func feature=gen-react type=test control=sequence
//ff:what writeComponentsUI 빈 컴포넌트 맵 시 fallback 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestWriteComponentsUI_EmptyComponents(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	spec := &design.DesignSpec{
		Components: map[string]design.ComponentToken{},
	}
	if err := writeComponentsUI(srcDir, spec); err != nil {
		t.Fatal(err)
	}
	uiDir := filepath.Join(srcDir, "components", "ui")
	if _, err := os.Stat(filepath.Join(uiDir, "Button.tsx")); os.IsNotExist(err) {
		t.Error("expected fallback Button.tsx for empty components map")
	}
}
