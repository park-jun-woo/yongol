//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what HasGuardCombinator — 조건 문자열에 결합/그룹 토큰(&& || 선행! 괄호) 존재 판별 검증

package stml

import "testing"

func TestHasGuardCombinator(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		want      bool
	}{
		{name: "and", condition: "a.b=x && c.d=y", want: true},
		{name: "or", condition: "a.b=x || c.d=y", want: true},
		{name: "open paren", condition: "(a.b=x", want: true},
		{name: "close paren", condition: "a.b=x)", want: true},
		{name: "leading bang", condition: "!a.b=x", want: true},
		{name: "leading bang with whitespace", condition: "  !a.b=x", want: true},
		{name: "single compare", condition: "a.b=x", want: false},
		{name: "lifecycle suffix", condition: "a.b.loading", want: false},
		{name: "not-equal is not combinator", condition: "a.b!=x", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasGuardCombinator(tt.condition)
			if got != tt.want {
				t.Errorf("HasGuardCombinator(%q) = %v, want %v", tt.condition, got, tt.want)
			}
		})
	}
}
