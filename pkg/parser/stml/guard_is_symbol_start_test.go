//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what isGuardSymbolStart — 토큰 시작 문자가 연산자/그룹 심볼인지 판별 검증

package stml

import "testing"

func TestIsGuardSymbolStart(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{name: "ampersand", r: '&', want: true},
		{name: "pipe", r: '|', want: true},
		{name: "bang", r: '!', want: true},
		{name: "open paren", r: '(', want: true},
		{name: "close paren", r: ')', want: true},
		{name: "equals", r: '=', want: true},
		{name: "less than", r: '<', want: true},
		{name: "greater than", r: '>', want: true},
		{name: "dot", r: '.', want: true},
		{name: "letter", r: 'a', want: false},
		{name: "digit", r: '1', want: false},
		{name: "space", r: ' ', want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGuardSymbolStart(tt.r)
			if got != tt.want {
				t.Errorf("isGuardSymbolStart(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
