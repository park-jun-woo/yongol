//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestScanCallInExprStmt — bare Scan ExprStmt 는 Assigned(errName="") 로 분류
package contract

import (
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
			assertScanCallInExprStmt(t, tt.body, tt.wantNil, tt.wantKind)
		})
	}
}
