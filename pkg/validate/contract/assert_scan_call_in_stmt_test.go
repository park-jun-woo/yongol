//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what assertScanCallInStmt — scanCallInStmt 디스패치 결과 검증 헬퍼
package contract

import "testing"

// assertScanCallInStmt parses body as a single statement and asserts the
// scanCallInStmt classification.
func assertScanCallInStmt(t *testing.T, body string, wantNil bool, wantKind scanKind, wantErr string) {
	t.Helper()
	stmt := mustFirstStmt(t, body)
	call, errName, kind := scanCallInStmt(stmt)
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
