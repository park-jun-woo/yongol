//ff:func feature=validate type=test control=iteration dimension=1 topic=funcspec-structural
//ff:what xff40FuncBodyTodo — HasBody=false인 func에 대해 XFF-40 진단 생성 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestXff40FuncBodyTodo(t *testing.T) {
	cases := []TestXff40FuncBodyTodoCase{
		{
			name:      "no_specs_returns_empty",
			specs:     nil,
			wantCount: 0,
		},
		{
			name: "has_body_no_diag",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", HasBody: true, Line: 5},
			},
			wantCount: 0,
		},
		{
			name: "stub_func_produces_diag",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", HasBody: false, Line: 10},
			},
			wantCount: 1,
		},
		{
			name: "mixed_body_and_stub",
			specs: []funcspec.FuncSpec{
				{Package: "billing", Name: "createInvoice", HasBody: true, Line: 5},
				{Package: "billing", Name: "deleteInvoice", HasBody: false, Line: 15},
				{Package: "auth", Name: "login", HasBody: false, Line: 20},
			},
			wantCount: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXff40FuncBodyTodo(t, c)
		})
	}
}
