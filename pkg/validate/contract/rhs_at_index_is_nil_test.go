//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestRhsAtIndexIsNil — AssignStmt RHS[idx] 가 literal nil 인지 판정 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestRhsAtIndexIsNil(t *testing.T) {
	tests := []struct {
		name string
		body string
		idx  int
		want bool
	}{
		{"err equals nil", "err = nil", 0, true},
		{"err equals call", "err = f()", 0, false},
		{"out of range high", "err = nil", 1, false},
		{"out of range low", "err = nil", -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			if got := rhsAtIndexIsNil(as, tt.idx); got != tt.want {
				t.Fatalf("rhsAtIndexIsNil(%q, %d) = %v, want %v", tt.body, tt.idx, got, tt.want)
			}
		})
	}
	if rhsAtIndexIsNil(nil, 0) {
		t.Fatal("nil assign should return false")
	}
}
