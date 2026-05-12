//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what optionalChainPath가 dotted path에 ?. 을 올바르게 삽입하는지 검증
package stml

import "testing"

func TestOptionalChainPath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"title", "title"},
		{"workflow.title", "workflow?.title"},
		{"summary.credits_balance", "summary?.credits_balance"},
		{"a.b.c", "a?.b?.c"},
		{"", ""},
	}
	for _, tt := range tests {
		got := optionalChainPath(tt.input)
		if got != tt.want {
			t.Errorf("optionalChainPath(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
