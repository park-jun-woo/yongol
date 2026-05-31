//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestCheckStyleColors — checkStyleColors 하드코딩 hex 색상 검출 분기 검증
package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCheckStyleColors(t *testing.T) {
	hexToToken := map[string]string{"#ff0000": "primary"}

	t.Run("no hex colors", func(t *testing.T) {
		var diags []diagnostic.Diagnostic
		checkStyleColors("display: flex", "Page.stml", hexToToken, &diags)
		if len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d", len(diags))
		}
	})

	t.Run("matching token fires once despite duplicate", func(t *testing.T) {
		var diags []diagnostic.Diagnostic
		checkStyleColors("color: #FF0000; border-color: #ff0000", "Page.stml", hexToToken, &diags)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XVM-05]") || !strings.Contains(diags[0].Message, "primary") {
			t.Errorf("unexpected message: %q", diags[0].Message)
		}
	})

	t.Run("hex with no matching token ignored", func(t *testing.T) {
		var diags []diagnostic.Diagnostic
		checkStyleColors("color: #00ff00", "Page.stml", hexToToken, &diags)
		if len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d", len(diags))
		}
	})
}
