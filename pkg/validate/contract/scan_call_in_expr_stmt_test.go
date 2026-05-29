//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanCallInExprStmt — bare Scan ExprStmt 는 Assigned(errName="") 로 분류

package contract

import (
	"go/ast"
	"testing"
)

func TestScanCallInExprStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind scanKind
	}{
		{"bare scan", "row.Scan(&x)", false, scanKindAssigned},
		{"non scan expr", "row.Next()", true, scanKindUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := mustFirstStmt(t, tt.body).(*ast.ExprStmt)
			call, errName, kind := scanCallInExprStmt(es)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
			if errName != "" {
				t.Errorf("errName = %q, want empty", errName)
			}
		})
	}
}
