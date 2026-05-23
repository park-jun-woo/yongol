//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what collectUnusedProps — 사용됨/미사용/ruleID별 advice 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectUnusedProps(t *testing.T) {
	fn := ssac.ServiceFunc{Name: "getUser", FileName: "user.ssac", Line: 5}
	used := map[string]bool{"id": true, "email": true}

	t.Run("all props used no diagnostics", func(t *testing.T) {
		diags := collectUnusedProps(fn, []string{"id", "email"}, used, "XOS-18")
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("unused prop raises warning", func(t *testing.T) {
		diags := collectUnusedProps(fn, []string{"phone"}, used, "XOS-18")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "XOS-18") {
			t.Errorf("Message missing rule: %s", diags[0].Message)
		}
	})

	t.Run("XSO-20 advice", func(t *testing.T) {
		diags := collectUnusedProps(fn, []string{"phone"}, used, "XSO-20")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Advice, "expose it from an SSaC variable") {
			t.Errorf("Advice missing expected text: %s", diags[0].Advice)
		}
	})
}
