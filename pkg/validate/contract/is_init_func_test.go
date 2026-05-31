//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIsInitFunc — FuncDecl 이 패키지 init() 함수인지 판정 검증
package contract

import (
	"testing"
)

func TestIsInitFunc(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want bool
	}{
		{"init func", "func init() {}", true},
		{"named func", "func setup() {}", false},
		{"init method", "func (r recv) init() {}", false},
		{"init with params", "func init() {}", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInitFunc(mustFuncDecl(t, tt.src)); got != tt.want {
				t.Fatalf("isInitFunc(%q) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
	if isInitFunc(nil) {
		t.Fatal("nil func decl should return false")
	}
}
