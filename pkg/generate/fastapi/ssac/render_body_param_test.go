//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderBodyParam — renderBodyParam body 파라미터 선언(없음/dict/Request 모델) 분기 검증
package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderBodyParam(t *testing.T) {
	plan := &ir.ServicePlan{OperationID: "CreateOrder"}
	cases := []struct {
		name         string
		hasBody      bool
		bodyFallback bool
		want         []string
	}{
		{"no body", false, false, nil},
		{"fallback dict", true, true, []string{"body: dict"}},
		{"request model", true, false, []string{"body: CreateOrderRequest"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := renderBodyParam(plan, c.hasBody, c.bodyFallback)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("renderBodyParam() = %v, want %v", got, c.want)
			}
		})
	}
}
