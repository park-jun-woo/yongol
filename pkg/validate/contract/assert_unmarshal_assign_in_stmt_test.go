//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertUnmarshalAssignInStmt — unmarshalAssignInStmt 디스패치 결과 검증 헬퍼
package contract

import "testing"

// assertUnmarshalAssignInStmt parses body as a single statement and asserts the
// unmarshalAssignInStmt classification.
func assertUnmarshalAssignInStmt(t *testing.T, body string, wantNil bool, wantKind unmarshalKind, wantErr string) {
	t.Helper()
	stmt := mustFirstStmt(t, body)
	_, call, errName, kind := unmarshalAssignInStmt(stmt)
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
