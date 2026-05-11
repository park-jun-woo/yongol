//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-06 테스트 — 중복 ## heading

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV06DuplicateHeading_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:     "DESIGN.md",
			Headings: []string{"Colors", "Typography", "Components"},
		},
	}
	if got := v06DuplicateHeading(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}

func TestV06DuplicateHeading_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:     "DESIGN.md",
			Headings: []string{"Colors", "Typography", "Colors"},
		},
	}
	got := v06DuplicateHeading(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[V-06]") {
		t.Fatalf("message missing [V-06] prefix: %q", got[0].Message)
	}
	if !strings.Contains(got[0].Message, "Colors") {
		t.Fatalf("message should mention the duplicate heading: %q", got[0].Message)
	}
}

func TestV06DuplicateHeading_Empty(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:     "DESIGN.md",
			Headings: nil,
		},
	}
	if got := v06DuplicateHeading(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
