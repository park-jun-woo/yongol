//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — design 토큰 참조/미지 prop 검사 헬퍼 직접 호출

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func designFS() *yongol.Fullstack {
	return &yongol.Fullstack{DesignSpec: &pdesign.DesignSpec{
		File:    "DESIGN.md",
		Colors:  map[string]string{"primary": "#fff"},
		Rounded: map[string]string{"md": "8px"},
		Spacing: map[string]string{"sm": "4px"},
	}}
}

func TestByNameResolveToken_ZeroCov(t *testing.T) {
	fs := designFS()
	if !resolveToken(fs, "colors.primary") {
		t.Errorf("colors.primary should resolve")
	}
	if !resolveToken(fs, "rounded.md") {
		t.Errorf("rounded.md should resolve")
	}
	if !resolveToken(fs, "spacing.sm") {
		t.Errorf("spacing.sm should resolve")
	}
	if resolveToken(fs, "colors.missing") {
		t.Errorf("missing color should not resolve")
	}
	if resolveToken(fs, "typography.body") {
		t.Errorf("typography branch nil map should not resolve")
	}
	if resolveToken(fs, "noseparator") {
		t.Errorf("no-dot ref should not resolve")
	}
	if resolveToken(fs, "unknown.group") {
		t.Errorf("unknown group should not resolve")
	}
}

func TestByNameCheckPropRefs_ZeroCov(t *testing.T) {
	fs := designFS()
	props := map[string]string{
		"color":  "{colors.primary}",   // resolves
		"border": "{colors.nope}",       // unresolved → diag
	}
	diags := checkPropRefs(fs, "Button", props)
	if len(diags) != 1 {
		t.Errorf("expected 1 unresolved-ref diag, got %d", len(diags))
	}

	// single prop helper directly: resolved → no diag.
	if d := checkSinglePropRefs(fs, "Button", "color", "{colors.primary}"); len(d) != 0 {
		t.Errorf("resolved single prop should yield no diag, got %v", d)
	}
	// no refs at all.
	if d := checkSinglePropRefs(fs, "Button", "label", "Save"); len(d) != 0 {
		t.Errorf("no-ref value should yield no diag")
	}
}

func TestByNameCheckUnknownProps_ZeroCov(t *testing.T) {
	props := map[string]string{
		"variant":    "primary", // known
		"weirdoProp": "x",        // unknown → warn
	}
	diags := checkUnknownProps("DESIGN.md", "Button", props)
	if len(diags) != 1 {
		t.Errorf("expected 1 unknown-prop warning, got %d", len(diags))
	}
}
