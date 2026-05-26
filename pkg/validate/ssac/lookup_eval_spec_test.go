//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what lookupEvalSpec — @eval pkg.Pascal → FuncSpec 검색 검증 (project 우선/builtin 대체/미매칭/빈 모델)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestLookupEvalSpec(t *testing.T) {
	project := []funcspec.FuncSpec{
		{Package: "billing", Name: "isZeroBalance"},
		{Package: "auth", Name: "hashPassword"},
	}
	builtin := []funcspec.FuncSpec{
		{Package: "billing", Name: "isZeroBalance"},
		{Package: "text", Name: "sanitize"},
	}
	cases := []struct {
		name, model, wantPkg, wantFn string
		wantNil                      bool
	}{
		{"project match", "billing.IsZeroBalance", "billing", "isZeroBalance", false},
		{"builtin fallback", "text.Sanitize", "text", "sanitize", false},
		{"no match", "billing.Unknown", "", "", true},
		{"no dot", "noDot", "", "", true},
		{"empty model", "", "", "", true},
		{"dot at start", ".Method", "", "", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := lookupEvalSpec(c.model, project, builtin)
			assertEvalSpec(t, c.model, got, c.wantNil, c.wantPkg, c.wantFn)
		})
	}
	t.Run("project pointer", func(t *testing.T) {
		got := lookupEvalSpec("auth.HashPassword", project, builtin)
		if got != &project[1] {
			t.Error("expected pointer into project slice")
		}
	})
	t.Run("builtin pointer", func(t *testing.T) {
		got := lookupEvalSpec("text.Sanitize", project, builtin)
		if got != &builtin[1] {
			t.Error("expected pointer into builtin slice")
		}
	})
}

