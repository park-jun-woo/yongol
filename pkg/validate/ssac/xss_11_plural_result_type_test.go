//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-11 — plural element type fires WARNING, singular passes, skip paths

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss11PluralResultType(t *testing.T) {
	t.Run("fires_on_plural", func(t *testing.T) {
		fs := xss11FS("get", "Gig.List", &parsessac.Result{Type: "Gigs", Var: "gigs", Wrapper: "Page"}, "")
		diags := xss11PluralResultType(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-11]") {
			t.Errorf("expected XSS-11 prefix, got %q", diags[0].Message)
		}
	})
	t.Run("passes_on_singular", func(t *testing.T) {
		diags := xss11PluralResultType(xss11FS("get", "Gig.List", &parsessac.Result{Type: "Gig", Var: "gigs", Wrapper: "Page"}, ""))
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})
	t.Run("skips_call", func(t *testing.T) {
		diags := xss11PluralResultType(xss11FS("call", "auth.Rotate", &parsessac.Result{Type: "Tokens", Var: "t"}, ""))
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})
	t.Run("skips_nil_result", func(t *testing.T) {
		diags := xss11PluralResultType(xss11FS("delete", "Item.Delete", nil, ""))
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})
	t.Run("skips_primitive", func(t *testing.T) {
		diags := xss11PluralResultType(xss11FS("get", "Item.Count", &parsessac.Result{Type: "int64", Var: "c"}, ""))
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})
	t.Run("skips_package_prefixed", func(t *testing.T) {
		diags := xss11PluralResultType(xss11FS("get", "Session.List", &parsessac.Result{Type: "Sessions", Var: "s"}, "session"))
		if len(diags) != 0 {
			t.Fatalf("expected 0 diags, got %d", len(diags))
		}
	})
}
