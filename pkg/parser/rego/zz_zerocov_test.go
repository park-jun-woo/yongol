//ff:func feature=policy type=test
//ff:what zz_zerocov_test — rego.collectOpaErrors 0% 커버리지 단위 테스트
package rego

import (
	"strings"
	"testing"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/ast/location"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCollectOpaErrors_ZeroCov(t *testing.T) {
	errs := ast.Errors{
		&ast.Error{Message: "with loc", Location: &location.Location{Row: 7}},
		&ast.Error{Message: "no loc"}, // Location nil → line 0
	}
	diags := collectOpaErrors("policy.rego", errs)
	if len(diags) != 2 {
		t.Fatalf("expected 2 diags, got %d", len(diags))
	}

	d0 := diags[0]
	if d0.File != "policy.rego" || d0.Line != 7 {
		t.Errorf("diag0 file/line = %q/%d", d0.File, d0.Line)
	}
	if d0.Phase != diagnostic.PhaseParse || d0.Level != diagnostic.LevelError {
		t.Errorf("diag0 phase/level = %v/%v", d0.Phase, d0.Level)
	}
	if !strings.Contains(d0.Message, "R-1") || !strings.Contains(d0.Message, "with loc") {
		t.Errorf("diag0 message = %q", d0.Message)
	}

	if diags[1].Line != 0 {
		t.Errorf("diag1 line = %d want 0", diags[1].Line)
	}

	// Empty error set → nil/empty slice.
	if got := collectOpaErrors("p.rego", ast.Errors{}); len(got) != 0 {
		t.Errorf("expected no diags for empty errors, got %d", len(got))
	}
}
