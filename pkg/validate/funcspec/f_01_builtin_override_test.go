//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what f01BuiltinOverride — 프로젝트 func가 built-in pkg func을 이름 충돌로 override할 때 WARNING 진단 생성 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestF01BuiltinOverride(t *testing.T) {
	cases := []TestF01BuiltinOverrideCase{
		{
			name:      "no_overlap_returns_empty",
			builtin:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}},
			project:   []funcspec.FuncSpec{{Package: "billing", Name: "createInvoice"}},
			wantCount: 0,
		},
		{
			name:      "same_pkg_different_name_no_diag",
			builtin:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}},
			project:   []funcspec.FuncSpec{{Package: "auth", Name: "verifyToken"}},
			wantCount: 0,
		},
		{
			name:      "different_pkg_same_name_no_diag",
			builtin:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}},
			project:   []funcspec.FuncSpec{{Package: "billing", Name: "hashPassword"}},
			wantCount: 0,
		},
		{
			name:      "exact_match_produces_warning",
			builtin:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}},
			project:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword", Line: 5}},
			wantCount: 1,
		},
		{
			name: "multiple_overlaps",
			builtin: []funcspec.FuncSpec{
				{Package: "auth", Name: "hashPassword"},
				{Package: "cache", Name: "get"},
			},
			project: []funcspec.FuncSpec{
				{Package: "auth", Name: "hashPassword", Line: 5},
				{Package: "cache", Name: "get", Line: 10},
				{Package: "billing", Name: "create", Line: 15},
			},
			wantCount: 2,
		},
		{
			name:      "empty_specs_returns_empty",
			builtin:   nil,
			project:   nil,
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runF01BuiltinOverride(t, c)
		})
	}
}
