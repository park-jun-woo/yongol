//ff:func feature=design-parse type=test control=sequence
//ff:what TestParseFile_Valid — DESIGN.md 파싱 정상 케이스 검증

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
