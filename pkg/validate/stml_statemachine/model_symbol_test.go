//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-statemachine
//ff:what modelSymbol — 가드 model 접두어를 StateDiagram.Symbol과 동일한 PascalCase로 정규화 검증

package stml_statemachine

import "testing"

func TestModelSymbol(t *testing.T) {
	tests := []struct {
		name  string
		model string
		want  string
	}{
		{name: "lowercase single word to PascalCase", model: "workflow", want: "Workflow"},
		{name: "already PascalCase is stable", model: "Workflow", want: "Workflow"},
		{name: "snake_case to PascalCase", model: "order_item", want: "OrderItem"},
		{name: "camelCase to PascalCase", model: "orderItem", want: "OrderItem"},
		{name: "empty string stays empty", model: "", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modelSymbol(tt.model); got != tt.want {
				t.Errorf("modelSymbol(%q) = %q, want %q", tt.model, got, tt.want)
			}
		})
	}
}
