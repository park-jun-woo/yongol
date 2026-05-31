//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertUnmarshalInAssignStmt — unmarshalInAssignStmt 결과 검증 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

// assertUnmarshalInAssignStmt parses body as an AssignStmt and asserts the
// unmarshalInAssignStmt classification.
func assertUnmarshalInAssignStmt(t *testing.T, body string, wantNil bool, wantKind unmarshalKind, wantErr string) {
	t.Helper()
	as := mustFirstStmt(t, body).(*ast.AssignStmt)
	_, call, errName, kind := unmarshalInAssignStmt(as)
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
