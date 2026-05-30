//ff:func feature=gen-react type=test control=sequence
//ff:what theme_picks_unit_test — pick_* 색상 접근자(via writeTailwindConfig) + orDefault + resolveTheme 단위 테스트
package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestWriteTailwindConfig_WithTheme exercises every pick_* accessor and the
// non-nil branch of orDefault by passing a fully populated FrontendTheme.
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

// TestWriteTailwindConfig_ThemeFallbackOnEmpty verifies orDefault falls back
// to the default when a theme field is empty even though theme is non-nil.
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

func TestResolveTheme(t *testing.T) {
	if got := resolveTheme(nil); got != nil {
		t.Errorf("nil fullstack -> %v, want nil", got)
	}
	if got := resolveTheme(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil manifest -> %v, want nil", got)
	}
	theme := &manifest.FrontendTheme{Primary: "#zzz"}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Theme: theme},
		},
	}
	if got := resolveTheme(fs); got != theme {
		t.Errorf("resolveTheme should return the embedded theme pointer, got %v", got)
	}
}
