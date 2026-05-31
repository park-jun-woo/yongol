//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestScanCallInIfStmt — if-init Scan 은 Discarded(가드됨) 로 분류
package contract

import (
	"go/ast"
	"testing"
)

func TestScanCallInIfStmt(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNil  bool
		wantKind scanKind
	}{
		{"if init scan", "if err := row.Scan(&x); err != nil { return }", false, scanKindDiscarded},
		{"if no init", "if x > 0 { return }", true, scanKindUnknown},
		{"if init non scan", "if err := row.Next(); err != nil { return }", true, scanKindUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifs := mustFirstStmt(t, tt.body).(*ast.IfStmt)
			call, _, kind := scanCallInIfStmt(ifs)
			if (call == nil) != tt.wantNil {
				t.Fatalf("call nil=%v, want %v", call == nil, tt.wantNil)
			}
			if kind != tt.wantKind {
				t.Errorf("kind = %v, want %v", kind, tt.wantKind)
			}
		})
	}
}
