//ff:func feature=gen-react type=test control=sequence
//ff:what writeTailwindConfig DESIGN.md 토큰 기반 설정 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestWriteTailwindConfig_WithDesignSpec(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	spec := &design.DesignSpec{
		Colors: map[string]string{
			"primary":    "#111111",
			"background": "#fafafa",
			"brand":      "#ff6600",
		},
		Rounded: map[string]string{
			"sm": "0.25rem",
			"md": "0.5rem",
			"lg": "0.75rem",
		},
		Spacing: map[string]string{
			"xs": "0.25rem",
			"sm": "0.5rem",
			"md": "1rem",
		},
	}
	if err := writeTailwindConfig(dir, nil, spec); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "tailwind.config.js"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	if !strings.Contains(content, "#111111") {
		t.Error("expected DesignSpec primary color #111111")
	}
	if !strings.Contains(content, "#fafafa") {
		t.Error("expected DesignSpec background color #fafafa")
	}
	if !strings.Contains(content, "'brand': '#ff6600'") {
		t.Error("expected extra DesignSpec color 'brand'")
	}
	if !strings.Contains(content, "sm: '0.25rem'") {
		t.Error("expected DesignSpec rounded sm")
	}
	if !strings.Contains(content, "lg: '0.75rem'") {
		t.Error("expected DesignSpec rounded lg")
	}
	if !strings.Contains(content, "spacing:") {
		t.Error("expected DesignSpec spacing section")
	}
	if !strings.Contains(content, "xs: '0.25rem'") {
		t.Error("expected DesignSpec spacing xs")
	}
}
