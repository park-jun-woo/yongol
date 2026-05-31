//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestScanCallInStmt — stmt 종류별 Scan 호출 분류 디스패치 검증
package contract

import (
	"testing"
)

func TestScanCallInStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind scanKind
		wantErr  string
	}{
		{"if init", "if err := row.Scan(&x); err != nil { return }", false, scanKindDiscarded, ""},
		{"assign", "err := row.Scan(&x)", false, scanKindAssigned, "err"},
		{"blank assign", "_ = row.Scan(&x)", false, scanKindDiscarded, ""},
		{"bare expr", "row.Scan(&x)", false, scanKindAssigned, ""},
		{"unrelated", "x := 1", true, scanKindUnknown, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertScanCallInStmt(t, tt.body, tt.wantNil, tt.wantKind, tt.wantErr)
		})
	}
}
