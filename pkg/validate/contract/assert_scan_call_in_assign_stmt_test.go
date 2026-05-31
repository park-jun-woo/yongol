//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertScanCallInAssignStmt — scanCallInAssignStmt 결과 검증 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

// assertScanCallInAssignStmt parses body as an AssignStmt and asserts the
// scanCallInAssignStmt classification.
func assertScanCallInAssignStmt(t *testing.T, body string, wantNil bool, wantKind scanKind, wantErr string) {
	t.Helper()
	as := mustFirstStmt(t, body).(*ast.AssignStmt)
	call, errName, kind := scanCallInAssignStmt(as)
	if (call == nil) != wantNil {
		t.Fatalf("call nil=%v, want %v", call == nil, wantNil)
	}
	if kind != wantKind {
		t.Errorf("kind = %v, want %v", kind, wantKind)
	}
	if errName != wantErr {
		t.Errorf("errName = %q, want %q", errName, wantErr)
	}
}
