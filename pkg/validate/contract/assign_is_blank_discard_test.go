//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestAssignIsBlankDiscard — AssignStmt LHS 가 모두 blank 인지 판정 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestAssignIsBlankDiscard(t *testing.T) {
	tests := []struct {
		name string
		body string
		want bool
	}{
		{"single blank", "_ = f()", true},
		{"all blank", "_, _ = f()", true},
		{"named lhs", "x := f()", false},
		{"mixed", "x, _ := f()", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			if got := assignIsBlankDiscard(as); got != tt.want {
				t.Fatalf("assignIsBlankDiscard(%q) = %v, want %v", tt.body, got, tt.want)
			}
		})
	}
	if assignIsBlankDiscard(nil) {
		t.Fatal("nil assign should return false")
	}
	if assignIsBlankDiscard(&ast.AssignStmt{}) {
		t.Fatal("empty lhs should return false")
	}
}
