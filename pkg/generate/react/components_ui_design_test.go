//ff:func feature=gen-react type=test control=sequence
//ff:what writeComponentsUI DESIGN.md 기반 컴포넌트 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestWriteComponentsUI_DesignComponents(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	spec := &design.DesignSpec{
		Components: map[string]design.ComponentToken{
			"Button": {
				Base: "inline-flex items-center justify-center",
				Variants: map[string]string{
					"primary":   "bg-primary",
					"secondary": "bg-secondary",
				},
				Sizes: map[string]string{
					"sm": "h-8",
					"md": "h-10",
				},
				DefaultVariant: "primary",
				DefaultSize:    "md",
			},
			"Card": {
				Base: "rounded-lg border",
			},
		},
	}
	if err := writeComponentsUI(srcDir, spec); err != nil {
		t.Fatal(err)
	}
	uiDir := filepath.Join(srcDir, "components", "ui")

	btnData, err := os.ReadFile(filepath.Join(uiDir, "Button.tsx"))
	if err != nil {
		t.Fatalf("Button.tsx not found: %v", err)
	}
	btnContent := string(btnData)
	if !strings.Contains(btnContent, "type Variant") {
		t.Error("Button.tsx should contain Variant type")
	}
	if !strings.Contains(btnContent, "type Size") {
		t.Error("Button.tsx should contain Size type")
	}

	cardData, err := os.ReadFile(filepath.Join(uiDir, "Card.tsx"))
	if err != nil {
		t.Fatalf("Card.tsx not found: %v", err)
	}
	cardContent := string(cardData)
	if !strings.Contains(cardContent, "rounded-lg border") {
		t.Error("Card.tsx should contain base class")
	}

	if _, err := os.Stat(filepath.Join(uiDir, "Modal.tsx")); !os.IsNotExist(err) {
		t.Error("Modal.tsx should not exist when design components are used")
	}
}
