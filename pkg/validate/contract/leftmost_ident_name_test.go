//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestLeftmostIdentName — 중첩 SelectorExpr 의 가장 왼쪽 Ident 이름 반환 검증

package contract

import "testing"

func TestLeftmostIdentName(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{"plain ident", "resp", "resp"},
		{"one selector", "resp.Body", "resp"},
		{"chained selector", "resp.Body.Reader", "resp"},
		{"call root", "f().Body", ""},
		{"literal root", "\"x\"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftmostIdentName(mustExpr(t, tt.src)); got != tt.want {
				t.Fatalf("leftmostIdentName(%q) = %q, want %q", tt.src, got, tt.want)
			}
		})
	}
}
