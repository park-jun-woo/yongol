//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHandlerParamDecls — writeHandlerParamDecls path/body/query 핸들러 파라미터 선언 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHandlerParamDecls(t *testing.T) {
	plan := &ir.ServicePlan{
		OperationID: "CreateOrder",
		PathParams:  []string{"id"},
		QueryParams: []ir.QueryParamMeta{{Name: "limit", Type: "integer", Required: true}},
	}

	t.Run("WithBody", func(t *testing.T) {
		var b strings.Builder
		writeHandlerParamDecls(&b, plan, true)
		got := b.String()
		for _, want := range []string{"id: int", "body: CreateOrderRequest", "limit: int"} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q in %q", want, got)
			}
		}
	})

	t.Run("WithoutBody", func(t *testing.T) {
		var b strings.Builder
		writeHandlerParamDecls(&b, plan, false)
		got := b.String()
		if strings.Contains(got, "body:") {
			t.Errorf("expected no body decl, got %q", got)
		}
		if !strings.Contains(got, "id: int") {
			t.Errorf("expected path decl, got %q", got)
		}
	})
}
