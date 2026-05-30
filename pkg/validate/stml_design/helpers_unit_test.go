//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestIsPureNumeric(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"4", true},
		{"0.5", true},
		{"1/2", true},
		{"", true}, // empty has no non-numeric chars
		{"sm", false},
		{"4px", false},
		{"-4", false},
	}
	for _, tt := range tests {
		if got := isPureNumeric(tt.in); got != tt.want {
			t.Errorf("isPureNumeric(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsOverrideComment(t *testing.T) {
	if !isOverrideComment(" @override class=\"x\" ") {
		t.Error("expected override comment recognised")
	}
	if isOverrideComment(" normal comment") {
		t.Error("expected non-override")
	}
}

func TestExtractOverrideClass(t *testing.T) {
	if got := extractOverrideClass(`@override class="card primary"`); got != "card primary" {
		t.Errorf("got %q, want 'card primary'", got)
	}
	if got := extractOverrideClass(`@override class='solo'`); got != "solo" {
		t.Errorf("single quote: %q", got)
	}
	if got := extractOverrideClass(`@override`); got != "" {
		t.Errorf("structure-only override should yield empty, got %q", got)
	}
}

func TestIsOverridden(t *testing.T) {
	ovr := overrideSet{
		"page.stml": {"card": true},
	}
	if !isOverridden(ovr, "page.stml", "card") {
		t.Error("expected card overridden")
	}
	if isOverridden(ovr, "page.stml", "button") {
		t.Error("button not overridden")
	}
	if isOverridden(ovr, "other.stml", "card") {
		t.Error("unknown file not overridden")
	}
	if isOverridden(ovr, "page.stml", "") {
		t.Error("empty class never overridden")
	}
}

func TestMatchColorPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchColorPrefix("bg-primary", "bg-primary", "p.stml", &out) {
		t.Fatal("expected bg-primary to match color")
	}
	if len(out.Colors) != 1 || out.Colors[0].Name != "primary" {
		t.Errorf("colors = %+v", out.Colors)
	}
	// Skippable name (numeric) → no record, no match.
	var out2 pageTokenRefs
	if matchColorPrefix("bg-500", "bg-500", "p.stml", &out2) {
		// matched? actually "500" -> isPureNumeric true -> continue -> returns false
		t.Error("numeric color value should be skippable")
	}
	// Non-color prefix → false.
	if matchColorPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex should not match a color prefix")
	}
}

func TestMatchFontPrefix(t *testing.T) {
	var out pageTokenRefs
	matchFontPrefix("font-display", "font-display", "p.stml", &out)
	if len(out.Fonts) != 1 || out.Fonts[0].Name != "display" {
		t.Errorf("fonts = %+v", out.Fonts)
	}
	// Skippable builtin → no record.
	var out2 pageTokenRefs
	matchFontPrefix("font-sans", "font-sans", "p.stml", &out2)
	if len(out2.Fonts) != 0 {
		t.Errorf("font-sans is builtin, should be skipped: %+v", out2.Fonts)
	}
	// Non-font prefix → no record.
	matchFontPrefix("text-lg", "text-lg", "p.stml", &out2)
	if len(out2.Fonts) != 0 {
		t.Error("text-lg should not record a font")
	}
}

func TestMatchRoundedPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchRoundedPrefix("rounded-card", "rounded-card", "p.stml", &out) {
		t.Fatal("expected rounded-card to match")
	}
	if len(out.Rounded) != 1 || out.Rounded[0].Name != "card" {
		t.Errorf("rounded = %+v", out.Rounded)
	}
	var out2 pageTokenRefs
	if matchRoundedPrefix("rounded-full", "rounded-full", "p.stml", &out2) {
		t.Error("rounded-full is builtin, should be skipped")
	}
	if matchRoundedPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex not a rounded prefix")
	}
}

func TestMatchSpacingPrefix(t *testing.T) {
	var out pageTokenRefs
	if !matchSpacingPrefix("p-section", "p-section", "p.stml", &out) {
		t.Fatal("expected p-section to match spacing")
	}
	if len(out.Spacing) != 1 || out.Spacing[0].Name != "section" {
		t.Errorf("spacing = %+v", out.Spacing)
	}
	var out2 pageTokenRefs
	if matchSpacingPrefix("p-4", "p-4", "p.stml", &out2) {
		t.Error("p-4 numeric should be skippable")
	}
	if matchSpacingPrefix("flex", "flex", "p.stml", &out2) {
		t.Error("flex not a spacing prefix")
	}
}

func TestClassifySingleToken(t *testing.T) {
	// Responsive prefix stripped, then color matched.
	var out pageTokenRefs
	classifySingleToken("sm:bg-primary", "sm:bg-primary", "p.stml", &out)
	if len(out.Colors) != 1 || out.Colors[0].Name != "primary" {
		t.Errorf("responsive color = %+v", out.Colors)
	}
	// Negative spacing prefix.
	var out2 pageTokenRefs
	classifySingleToken("-mt-gutter", "-mt-gutter", "p.stml", &out2)
	if len(out2.Spacing) != 1 || out2.Spacing[0].Name != "gutter" {
		t.Errorf("negative spacing = %+v", out2.Spacing)
	}
	// Font fallback.
	var out3 pageTokenRefs
	classifySingleToken("font-brand", "font-brand", "p.stml", &out3)
	if len(out3.Fonts) != 1 {
		t.Errorf("font = %+v", out3.Fonts)
	}
}

func TestClassifyTokens(t *testing.T) {
	var out pageTokenRefs
	classifyTokens("bg-primary rounded-card p-section", "p.stml", &out)
	if len(out.Colors) != 1 || len(out.Rounded) != 1 || len(out.Spacing) != 1 {
		t.Errorf("classifyTokens result: %+v", out)
	}
	// Empty class is a no-op.
	var out2 pageTokenRefs
	classifyTokens("", "p.stml", &out2)
	if len(out2.Colors)+len(out2.Spacing)+len(out2.Rounded)+len(out2.Fonts) != 0 {
		t.Error("empty class should produce nothing")
	}
}

func TestIsSkippable(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"[10px]", true},
		{"gray-500", true},
		{"full", true},
		{"none", true},
		{"semibold", true},
		{"4", true},
		{"0.5", true},
		{"primary", false},
		{"display", false},
	}
	for _, tt := range tests {
		if got := isSkippable(tt.in); got != tt.want {
			t.Errorf("isSkippable(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]string{"c": "1", "a": "2", "b": "3"}
	if got := sortedKeys(m); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedKeys = %v", got)
	}
	if got := sortedKeys(map[string]string{}); len(got) != 0 {
		t.Errorf("empty map should give empty slice, got %v", got)
	}
}

func TestSortedCompKeys(t *testing.T) {
	m := map[string]design.ComponentToken{"z": {}, "a": {}}
	if got := sortedCompKeys(m); !reflect.DeepEqual(got, []string{"a", "z"}) {
		t.Errorf("sortedCompKeys = %v", got)
	}
}

func TestSortedTypoKeys(t *testing.T) {
	m := map[string]design.TypographyToken{"y": {}, "b": {}}
	if got := sortedTypoKeys(m); !reflect.DeepEqual(got, []string{"b", "y"}) {
		t.Errorf("sortedTypoKeys = %v", got)
	}
}
