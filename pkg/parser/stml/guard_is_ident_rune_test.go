//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what isGuardIdentRune — 식별자/숫자/enum-literal 내부 문자 판별 검증

package stml

import "testing"

func TestIsGuardIdentRune(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{name: "lower letter", r: 'a', want: true},
		{name: "upper letter", r: 'Z', want: true},
		{name: "digit", r: '5', want: true},
		{name: "underscore", r: '_', want: true},
		{name: "hyphen", r: '-', want: true},
		{name: "dot", r: '.', want: false},
		{name: "space", r: ' ', want: false},
		{name: "equals", r: '=', want: false},
		{name: "unicode letter", r: '한', want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGuardIdentRune(tt.r)
			if got != tt.want {
				t.Errorf("isGuardIdentRune(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
