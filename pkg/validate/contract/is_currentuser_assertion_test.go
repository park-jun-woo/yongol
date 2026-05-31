//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIsCurrentUserAssertion — expr 가 ctx.Value("currentUser").(T) 단언인지 검증
package contract

import (
	"testing"
)

func TestIsCurrentUserAssertion(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"valid", "ctx.Value(\"currentUser\").(User)", true},
		{"wrong key", "ctx.Value(\"other\").(User)", false},
		{"wrong receiver", "c.Value(\"currentUser\").(User)", false},
		{"wrong method", "ctx.Get(\"currentUser\").(User)", false},
		{"not assertion", "ctx.Value(\"currentUser\")", false},
		{"non literal key", "ctx.Value(k).(User)", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCurrentUserAssertion(mustExpr(t, tt.src)); got != tt.want {
				t.Fatalf("isCurrentUserAssertion(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
}
