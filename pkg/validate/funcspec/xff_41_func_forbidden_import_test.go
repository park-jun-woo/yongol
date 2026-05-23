//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what xff41FuncForbiddenImport — ProjectFuncSpecs의 forbidden import에 대해 XFF-41 진단 생성 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestXff41FuncForbiddenImport(t *testing.T) {
	cases := []TestXff41FuncForbiddenImportCase{
		{
			name:      "no_specs_returns_empty",
			specs:     nil,
			wantCount: 0,
		},
		{
			name: "allowed_imports_no_diag",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", Imports: []string{"fmt", "strings"}, Line: 5},
			},
			wantCount: 0,
		},
		{
			name: "forbidden_import_produces_diag",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", Imports: []string{"database/sql"}, Line: 5},
			},
			wantCount: 1,
		},
		{
			name: "multiple_forbidden_imports_in_one_func",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", Imports: []string{"database/sql", "net/http"}, Line: 5},
			},
			wantCount: 2,
		},
		{
			name: "multiple_funcs_with_forbidden_imports",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", Imports: []string{"database/sql"}, Line: 5},
				{Package: "auth", Name: "login", Imports: []string{"net/http"}, Line: 10},
			},
			wantCount: 2,
		},
		{
			name: "no_imports_field_no_diag",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", Line: 5},
			},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXff41FuncForbiddenImport(t, c)
		})
	}
}
