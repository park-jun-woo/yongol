//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestHandlerQueryParamDecl — handlerQueryParamDecl required/optional 핸들러 시그니처 선언 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHandlerQueryParamDecl(t *testing.T) {
	cases := []struct {
		name string
		qp   ir.QueryParamMeta
		want string
	}{
		{
			name: "required int",
			qp:   ir.QueryParamMeta{Name: "limit", Type: "integer", Required: true},
			want: "limit: int",
		},
		{
			name: "optional str",
			qp:   ir.QueryParamMeta{Name: "cursor", Type: "string", Required: false},
			want: "cursor: str | None = None",
		},
		{
			name: "optional bool",
			qp:   ir.QueryParamMeta{Name: "active", Type: "boolean", Required: false},
			want: "active: bool | None = None",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := handlerQueryParamDecl(c.qp); got != c.want {
				t.Errorf("handlerQueryParamDecl(%+v) = %q, want %q", c.qp, got, c.want)
			}
		})
	}
}
