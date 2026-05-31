//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestUnmarshalErrName — AssignStmt LHS 에서 error 식별자 이름 추출 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestUnmarshalErrName(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{"err", "err := json.Unmarshal(b, &v)", "err"},
		{"suffix err", "uErr := json.Unmarshal(b, &v)", "uErr"},
		{"blank discard", "_ = json.Unmarshal(b, &v)", ""},
		{"non err name", "x := json.Unmarshal(b, &v)", ""},
		{"multi lhs", "a, b := f()", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			if got := unmarshalErrName(as); got != tt.want {
				t.Fatalf("unmarshalErrName(%q) = %q, want %q", tt.body, got, tt.want)
			}
		})
	}
	if unmarshalErrName(nil) != "" {
		t.Fatal("nil assign should return empty")
	}
}
