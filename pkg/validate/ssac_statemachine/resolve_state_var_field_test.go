//ff:func feature=validate type=test control=selection topic=ssac-statemachine
//ff:what TestResolveStateVarField — resolveStateVarField var.Field 타입 체인 해석 분기 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestResolveStateVarField(t *testing.T) {
	g := &rule.Ground{Types: map[string]string{
		"SSaC.var.Fn.order":      "[]*Order",
		"DDL.field.Order.status": "string",
	}}

	t.Run("unknown var returns empty", func(t *testing.T) {
		if got := resolveStateVarField(g, "Fn", "missing", "status"); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("resolved field type with slice and pointer trimmed", func(t *testing.T) {
		if got := resolveStateVarField(g, "Fn", "order", "status"); got != "string" {
			t.Errorf("expected string, got %q", got)
		}
	})
}
