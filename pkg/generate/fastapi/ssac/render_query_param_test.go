//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderQueryParam — renderQueryParam required/optional Python query 파라미터 선언 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderQueryParam(t *testing.T) {
	cases := []struct {
		name string
		qp   ir.QueryParamMeta
		want string
	}{
		{"required number", ir.QueryParamMeta{Name: "ratio", Type: "number", Required: true}, "ratio: float"},
		{"optional string", ir.QueryParamMeta{Name: "q", Type: "string", Required: false}, "q: str | None = None"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := renderQueryParam(c.qp); got != c.want {
				t.Errorf("renderQueryParam(%+v) = %q, want %q", c.qp, got, c.want)
			}
		})
	}
}
