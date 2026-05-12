//ff:func feature=frontend type=test
//ff:what ParseFile — DESIGN.md 파싱 정상 케이스 및 에러 케이스 검증
package design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_Valid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	content := `---
version: "1.0"
name: TestDesign
colors:
  primary: "#111111"
  secondary: "#222222"
typography:
  heading:
    fontFamily: "Inter"
    fontSize: "2rem"
    fontWeight: "700"
    lineHeight: "1.2"
    letterSpacing: "-0.01em"
rounded:
  sm: "4px"
spacing:
  xs: "4px"
  md: "16px"
components:
  Card:
    props:
      elevation: "1 | 2 | 3"
---

## Colors

Description of colors.

## Typography

Description of typography.
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
	if spec.File != path {
		t.Errorf("File: want %q, got %q", path, spec.File)
	}
	if spec.Version != "1.0" {
		t.Errorf("Version: want %q, got %q", "1.0", spec.Version)
	}
	if spec.Name != "TestDesign" {
		t.Errorf("Name: want %q, got %q", "TestDesign", spec.Name)
	}
	if len(spec.Colors) != 2 {
		t.Errorf("Colors count: want 2, got %d", len(spec.Colors))
	}
	if spec.Colors["primary"] != "#111111" {
		t.Errorf("Colors[primary]: want %q, got %q", "#111111", spec.Colors["primary"])
	}
	if len(spec.Typography) != 1 {
		t.Errorf("Typography count: want 1, got %d", len(spec.Typography))
	}
	if spec.Typography["heading"].FontFamily != "Inter" {
		t.Errorf("Typography[heading].FontFamily: want %q, got %q", "Inter", spec.Typography["heading"].FontFamily)
	}
	if spec.Typography["heading"].FontSize != "2rem" {
		t.Errorf("Typography[heading].FontSize: want %q, got %q", "2rem", spec.Typography["heading"].FontSize)
	}
	if len(spec.Rounded) != 1 {
		t.Errorf("Rounded count: want 1, got %d", len(spec.Rounded))
	}
	if len(spec.Spacing) != 2 {
		t.Errorf("Spacing count: want 2, got %d", len(spec.Spacing))
	}
	if len(spec.Components) != 1 {
		t.Errorf("Components count: want 1, got %d", len(spec.Components))
	}
	if spec.Components["Card"].Props["elevation"] != "1 | 2 | 3" {
		t.Errorf("Components[Card].Props[elevation]: want %q, got %q", "1 | 2 | 3", spec.Components["Card"].Props["elevation"])
	}
	if len(spec.Headings) != 2 {
		t.Errorf("Headings count: want 2, got %d", len(spec.Headings))
	}
	if spec.Headings[0] != "Colors" {
		t.Errorf("Headings[0]: want %q, got %q", "Colors", spec.Headings[0])
	}
	if spec.Headings[1] != "Typography" {
		t.Errorf("Headings[1]: want %q, got %q", "Typography", spec.Headings[1])
	}
}

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

func TestParseFile_MissingFile(t *testing.T) {
	_, diags := ParseFile("/nonexistent/DESIGN.md")
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("expected ERROR level, got %q", diags[0].Level)
	}
}

func TestParseFile_NoFrontMatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	if err := os.WriteFile(path, []byte("# No front matter\n"), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseFile(path)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("expected ERROR level, got %q", diags[0].Level)
	}
}

func TestParseFile_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	content := "---\n: invalid: yaml: [broken\n---\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseFile(path)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("expected ERROR level, got %q", diags[0].Level)
	}
}
