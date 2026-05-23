//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what Run — FuncSpec 검증 전체 실행 (F-1, XFF-40, XFF-41) 통합 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	cases := []TestRunCase{
		{
			name:      "empty_fullstack_no_diags",
			fs:        &yongol.Fullstack{},
			wantCount: 0,
		},
		{
			name: "clean_project_func_no_diags",
			fs: &yongol.Fullstack{
				ProjectFuncSpecs: []funcspec.FuncSpec{
					{Package: "billing", Name: "createInvoice", HasBody: true, Imports: []string{"fmt"}},
				},
			},
			wantCount: 0,
		},
		{
			name: "stub_func_triggers_xff40",
			fs: &yongol.Fullstack{
				ProjectFuncSpecs: []funcspec.FuncSpec{
					{Package: "billing", Name: "createInvoice", HasBody: false, Line: 5},
				},
			},
			wantCount: 1,
			wantCodes: []string{"[XFF-40]"},
		},
		{
			name: "forbidden_import_triggers_xff41",
			fs: &yongol.Fullstack{
				ProjectFuncSpecs: []funcspec.FuncSpec{
					{Package: "billing", Name: "createInvoice", HasBody: true, Imports: []string{"database/sql"}, Line: 5},
				},
			},
			wantCount: 1,
			wantCodes: []string{"[XFF-41]"},
		},
		{
			name: "builtin_override_triggers_f1",
			fs: &yongol.Fullstack{
				YongolPkgSpecs:   []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}},
				ProjectFuncSpecs: []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword", HasBody: true, Line: 3}},
			},
			wantCount: 1,
			wantCodes: []string{"[F-1]"},
		},
		{
			name: "all_three_rules_fire",
			fs: &yongol.Fullstack{
				YongolPkgSpecs: []funcspec.FuncSpec{{Package: "auth", Name: "hash"}},
				ProjectFuncSpecs: []funcspec.FuncSpec{
					{Package: "auth", Name: "hash", HasBody: false, Imports: []string{"net/http"}, Line: 1},
				},
			},
			wantCount: 3,
			wantCodes: []string{"[F-1]", "[XFF-40]", "[XFF-41]"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runRun(t, c)
		})
	}
}
