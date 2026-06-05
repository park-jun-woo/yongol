//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TM-17 — 가드 조건 진단: 결합 토큰 유효/무효 + 레거시 조건 스킵

package stml_openapi

import "testing"

func TestTM17GuardSyntax(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		wantDiag  bool
	}{
		{"valid and guard", "workflow.status=active && currentUser.Role=owner", false},
		{"valid group", "(workflow.status=a || workflow.status=b)", false},
		{"valid negation", "!workflow.locked=true", false},
		{"invalid unclosed group", "(workflow.status=a", true},
		{"invalid dangling combinator", "workflow.status=a &&", true},
		{"invalid arithmetic in guard", "workflow.status=a && b.y=1+2", true},
		{"legacy field=value skipped", "workflow.status=draft", false},
		{"legacy loading skipped", ".loading", false},
		{"legacy empty skipped", "items.empty", false},
		{"legacy bare field skipped", "items", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTM17(t, tt.condition, tt.wantDiag)
		})
	}
}
