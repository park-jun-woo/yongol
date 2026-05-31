//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestCallMethodName — CallExpr 가 SelectorExpr 면 method 이름 반환 검증
package contract

import (
	"testing"
)

func TestCallMethodName(t *testing.T) {
	tests := []struct {
		name string
		src  string
		want string
	}{
		{"method call", "x.Close()", "Close"},
		{"chained method", "resp.Body.Close()", "Close"},
		{"free function", "doThing()", ""},
		{"pkg func", "fmt.Sprintf(\"x\")", "Sprintf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callMethodName(mustCall(t, tt.src)); got != tt.want {
				t.Fatalf("callMethodName(%q) = %q, want %q", tt.src, got, tt.want)
			}
		})
	}
	if callMethodName(nil) != "" {
		t.Fatal("nil call should return empty string")
	}
}
