//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV07UnknownProp_Negative — 미지 prop 시 V-07 WARNING 진단 1건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV07UnknownProp_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Card": {Props: map[string]string{
					"variant":     "solid",
					"customField": "xyz",
				}},
			},
		},
	}
	got := v07UnknownProp(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[V-07]") {
		t.Fatalf("message missing [V-07] prefix: %q", got[0].Message)
	}
	if got[0].Level != diagnostic.LevelWarning {
		t.Fatalf("V-07 should be WARNING, got %q", got[0].Level)
	}
}
