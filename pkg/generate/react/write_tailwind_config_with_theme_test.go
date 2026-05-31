//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what theme_picks_unit_test — pick_* 색상 접근자(via writeTailwindConfig) + orDefault + resolveTheme 단위 테스트
package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestWriteTailwindConfig_WithTheme(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "src"), 0o755); err != nil {
		t.Fatal(err)
	}
	theme := &manifest.FrontendTheme{
		Primary:     "#111111",
		Secondary:   "#222222",
		Accent:      "#333333",
		Destructive: "#444444",
		Muted:       "#555555",
		Background:  "#666666",
		Foreground:  "#777777",
		Border:      "#888888",
		Radius:      "0.75rem",
	}
	if err := writeTailwindConfig(dir, theme, nil); err != nil {
		t.Fatalf("writeTailwindConfig: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "tailwind.config.js"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	for name, color := range map[string]string{
		"primary":     "#111111",
		"secondary":   "#222222",
		"accent":      "#333333",
		"destructive": "#444444",
		"muted":       "#555555",
		"background":  "#666666",
		"foreground":  "#777777",
		"border":      "#888888",
	} {
		if !strings.Contains(content, color) {
			t.Errorf("theme color for %s (%s) not present in tailwind config", name, color)
		}
	}
}
