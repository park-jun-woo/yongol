//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanCallFromAssign — AssignStmt RHS 가 Scan 호출이면 CallExpr 반환 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestScanCallFromAssign(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantNil bool
	}{
		{"scan rhs", "err := row.Scan(&x)", false},
		{"two value scan", "err = row.Scan(&x)", false},
		{"non scan call", "err := row.Next()", true},
		{"not a call", "x := 1", true},
		{"multi rhs", "a, b := f(), g()", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as, ok := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			if !ok {
				t.Fatalf("body %q is not an AssignStmt", tt.body)
			}
			got := scanCallFromAssign(as)
			if (got == nil) != tt.wantNil {
				t.Fatalf("scanCallFromAssign(%q) nil=%v, want nil=%v", tt.body, got == nil, tt.wantNil)
			}
		})
	}
	if scanCallFromAssign(nil) != nil {
		t.Fatal("nil assign should return nil")
	}
}
