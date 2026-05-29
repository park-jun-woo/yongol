//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanCallInAssignStmt — AssignStmt Scan 호출의 kind/errName 분류 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestScanCallInAssignStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind scanKind
		wantErr  string
	}{
		{"assigned err", "err := row.Scan(&x)", false, scanKindAssigned, "err"},
		{"blank discard", "_ = row.Scan(&x)", false, scanKindDiscarded, ""},
		{"not scan", "err := row.Next()", true, scanKindUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			call, errName, kind := scanCallInAssignStmt(as)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
			if errName != tt.wantErr {
				t.Errorf("errName = %q, want %q", errName, tt.wantErr)
			}
		})
	}
}
