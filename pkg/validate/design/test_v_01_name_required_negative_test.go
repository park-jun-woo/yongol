//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV01NameRequired_Negative — name 필드 누락 시 V-01 진단 1건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV01NameRequired_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Name: "",
		},
	}
	got := v01NameRequired(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[V-01]") {
		t.Fatalf("message missing [V-01] prefix: %q", got[0].Message)
	}
}
