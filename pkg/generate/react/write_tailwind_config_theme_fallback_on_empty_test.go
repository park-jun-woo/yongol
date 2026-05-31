//ff:func feature=gen-react type=test control=sequence
//ff:what theme_picks_unit_test — pick_* 색상 접근자(via writeTailwindConfig) + orDefault + resolveTheme 단위 테스트
package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestWriteTailwindConfig_ThemeFallbackOnEmpty(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "src"), 0o755); err != nil {
		t.Fatal(err)
	}
	// Only Primary set; the rest are empty -> defaults used.
	theme := &manifest.FrontendTheme{Primary: "#abcabc"}
	if err := writeTailwindConfig(dir, theme, nil); err != nil {
		t.Fatalf("writeTailwindConfig: %v", err)
	}
	data, _ := os.ReadFile(filepath.Join(dir, "tailwind.config.js"))
	content := string(data)
	if !strings.Contains(content, "#abcabc") {
		t.Errorf("set primary color missing: %s", content)
	}
	// default background should appear when theme.Background == ""
	if !strings.Contains(content, "#ffffff") {
		t.Errorf("expected default background #ffffff when theme.Background empty")
	}
}
