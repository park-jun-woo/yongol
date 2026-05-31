//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertScanCallInExprStmt — scanCallInExprStmt 결과(call nil / kind / errName) 검증 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

// assertScanCallInExprStmt parses body as a single ExprStmt and asserts
// scanCallInExprStmt's call-nil-ness, kind, and empty errName.
func assertScanCallInExprStmt(t *testing.T, body string, wantNil bool, wantKind scanKind) {
	t.Helper()
	es := mustFirstStmt(t, body).(*ast.ExprStmt)
	call, errName, kind := scanCallInExprStmt(es)
	if (call == nil) != wantNil {
		t.Fatalf("call nil=%v, want %v", call == nil, wantNil)
	}
	if kind != wantKind {
		t.Errorf("kind = %v, want %v", kind, wantKind)
	}
	if errName != "" {
		t.Errorf("errName = %q, want empty", errName)
	}
}
