//ff:func feature=design-parse type=test control=sequence
//ff:what TestParseFile_ComponentTokenExtended — components YAML 확장 필드 파싱 검증

package design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_ComponentTokenExtended(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	content := `---
version: "1.0"
name: ExtendedTest
components:
  Button:
    base: "inline-flex items-center justify-center"
    variants:
      primary: "bg-primary text-white"
      secondary: "bg-secondary text-black"
    sizes:
      sm: "h-8 px-3"
      md: "h-10 px-4"
    defaultVariant: primary
    defaultSize: md
    props:
      icon: "ReactNode"
  Card:
    base: "rounded-lg border shadow-sm"
---

## Components
`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("spec is nil")
	}
	if len(spec.Components) != 2 {
		t.Fatalf("Components count: want 2, got %d", len(spec.Components))
	}

	btn := spec.Components["Button"]
	if btn.Base != "inline-flex items-center justify-center" {
		t.Errorf("Button.Base: want %q, got %q", "inline-flex items-center justify-center", btn.Base)
	}
	if len(btn.Variants) != 2 {
		t.Errorf("Button.Variants count: want 2, got %d", len(btn.Variants))
	}
	if btn.Variants["primary"] != "bg-primary text-white" {
		t.Errorf("Button.Variants[primary]: want %q, got %q", "bg-primary text-white", btn.Variants["primary"])
	}
	if len(btn.Sizes) != 2 {
		t.Errorf("Button.Sizes count: want 2, got %d", len(btn.Sizes))
	}
	if btn.Sizes["sm"] != "h-8 px-3" {
		t.Errorf("Button.Sizes[sm]: want %q, got %q", "h-8 px-3", btn.Sizes["sm"])
	}
	if btn.DefaultVariant != "primary" {
		t.Errorf("Button.DefaultVariant: want %q, got %q", "primary", btn.DefaultVariant)
	}
	if btn.DefaultSize != "md" {
		t.Errorf("Button.DefaultSize: want %q, got %q", "md", btn.DefaultSize)
	}
	if btn.Props["icon"] != "ReactNode" {
		t.Errorf("Button.Props[icon]: want %q, got %q", "ReactNode", btn.Props["icon"])
	}

	card := spec.Components["Card"]
	if card.Base != "rounded-lg border shadow-sm" {
		t.Errorf("Card.Base: want %q, got %q", "rounded-lg border shadow-sm", card.Base)
	}
	if len(card.Variants) != 0 {
		t.Errorf("Card.Variants count: want 0, got %d", len(card.Variants))
	}
	if len(card.Sizes) != 0 {
		t.Errorf("Card.Sizes count: want 0, got %d", len(card.Sizes))
	}
}
