//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-statemachine
//ff:what TestResolveStateInputType — resolveStateInputType @state input 표현식 타입 해석 분기 검증
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestResolveStateInputType(t *testing.T) {
	g := &rule.Ground{Types: map[string]string{
		"Manifest.claim.role":    "string",
		"SSaC.var.Fn.order":      "*Order",
		"DDL.field.Order.status": "string",
		"SSaC.var.Fn.plainVar":   "int",
	}}

	cases := []struct {
		name  string
		value string
		want  string
	}{
		{name: "empty", value: "  ", want: ""},
		{name: "literal int", value: "42", want: "int"},
		{name: "currentUser claim", value: "currentUser.role", want: "string"},
		{name: "var.field", value: "order.status", want: "string"},
		{name: "contains special char only", value: ".foo", want: ""},
		{name: "plain var", value: "plainVar", want: "int"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := resolveStateInputType(g, "Fn", tc.value); got != tc.want {
				t.Errorf("resolveStateInputType(%q) = %q, want %q", tc.value, got, tc.want)
			}
		})
	}
}
