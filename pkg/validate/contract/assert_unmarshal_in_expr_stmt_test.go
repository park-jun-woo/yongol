//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertUnmarshalInExprStmt — unmarshalInExprStmt 결과(errName="") 검증 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

// assertUnmarshalInExprStmt parses body as an ExprStmt and asserts the
// unmarshalInExprStmt classification with empty errName.
func assertUnmarshalInExprStmt(t *testing.T, body string, wantNil bool, wantKind unmarshalKind) {
	t.Helper()
	es := mustFirstStmt(t, body).(*ast.ExprStmt)
	_, call, errName, kind := unmarshalInExprStmt(es)
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
