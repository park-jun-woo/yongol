//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestGuardCollectRefs — 가드 AST에서 model.Field ref를 좌→우 순서로 수집

package stml

import "testing"

func TestGuardCollectRefs(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		want      []string
	}{
		{"single compare", "workflow.status = draft", []string{"workflow.status"}},
		{"and combination", "workflow.status=active && currentUser.Role=owner", []string{"workflow.status", "currentUser.Role"}},
		{"negation and group", "!(workflow.locked = true)", []string{"workflow.locked"}},
		{"lifecycle", "items.list.loading", []string{"items.list"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertGuardCollectRefs(t, tt.condition, tt.want)
		})
	}
}
