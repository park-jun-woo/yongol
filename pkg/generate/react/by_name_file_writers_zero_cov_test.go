//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestByNameFileWriters_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeLibUtils(dir); err != nil {
		t.Fatalf("writeLibUtils: %v", err)
	}
	if err := writeAppTSXPlaceholder(dir); err != nil {
		t.Fatalf("writeAppTSXPlaceholder: %v", err)
	}

	uiDir := filepath.Join(dir, "ui")
	if err := os.MkdirAll(uiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	comps := map[string]design.ComponentToken{
		"Input":  {Base: "px-2"},
		"Button": {Variants: map[string]string{"primary": "bg-blue"}},
	}
	if err := writeDesignComponents(uiDir, comps); err != nil {
		t.Fatalf("writeDesignComponents: %v", err)
	}
}
