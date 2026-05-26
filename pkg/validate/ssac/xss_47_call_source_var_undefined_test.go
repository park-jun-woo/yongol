//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what XSS-47 — undefined @call source var fires, declared/implicit/empty/non-call pass

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXss47CallSourceVarUndefined(t *testing.T) {
	t.Run("fires_on_undefined", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Bad", FileName: "s/bad.ssac",
			Sequences: []parsessac.Sequence{{Type: "call", Model: "billing.Charge", Args: []parsessac.Arg{{Source: "order", Field: "ID"}}, Line: 3}},
		}}}
		diags := xss47CallSourceVarUndefined(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XSS-47]") {
			t.Errorf("prefix mismatch: %q", diags[0].Message)
		}
	})
	t.Run("passes_on_declared", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Good", FileName: "s/good.ssac",
			Sequences: []parsessac.Sequence{
				{Type: "get", Model: "Order.FindByID", Result: &parsessac.Result{Var: "order", Type: "Order"}, Line: 3},
				{Type: "call", Model: "billing.Charge", Args: []parsessac.Arg{{Source: "order", Field: "ID"}}, Line: 5},
			},
		}}}
		diags := xss47CallSourceVarUndefined(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("passes_on_implicit", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Impl", FileName: "s/impl.ssac",
			Sequences: []parsessac.Sequence{{Type: "call", Model: "auth.Login", Args: []parsessac.Arg{{Source: "request", Field: "Body"}}, Line: 3}},
		}}}
		diags := xss47CallSourceVarUndefined(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_non_call", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Get", FileName: "s/get.ssac",
			Sequences: []parsessac.Sequence{{Type: "get", Model: "C.FindByID", Args: []parsessac.Arg{{Source: "undefined", Field: "ID"}}, Line: 3}},
		}}}
		diags := xss47CallSourceVarUndefined(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("passes_on_empty_source", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{
			Name: "Lit", FileName: "s/lit.ssac",
			Sequences: []parsessac.Sequence{{Type: "call", Model: "billing.Charge", Args: []parsessac.Arg{{Source: "", Literal: "100"}}, Line: 3}},
		}}}
		diags := xss47CallSourceVarUndefined(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}
