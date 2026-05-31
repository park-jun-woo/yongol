//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestScanCallInAssignStmt — AssignStmt Scan 호출의 kind/errName 분류 검증
package contract

import (
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
			assertScanCallInAssignStmt(t, tt.body, tt.wantNil, tt.wantKind, tt.wantErr)
		})
	}
}
