//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-statemachine
//ff:what transitionFrom — diagram에서 op 이벤트의 From 상태 반환 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestTransitionFrom(t *testing.T) {
	d := &statemachine.StateDiagram{
		Transitions: []statemachine.Transition{
			{From: "draft", To: "submitted", Event: "submitOrder"},
			{From: "submitted", To: "approved", Event: "approveOrder"},
		},
	}

	cases := []struct {
		name string
		d    *statemachine.StateDiagram
		op   string
		want string
	}{
		{name: "nil_diagram", d: nil, op: "submitOrder", want: ""},
		{name: "found", d: d, op: "submitOrder", want: "draft"},
		{name: "found_second", d: d, op: "approveOrder", want: "submitted"},
		{name: "not_found", d: d, op: "deleteOrder", want: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := transitionFrom(c.d, c.op)
			if got != c.want {
				t.Errorf("transitionFrom(..., %q) = %q, want %q", c.op, got, c.want)
			}
		})
	}
}
