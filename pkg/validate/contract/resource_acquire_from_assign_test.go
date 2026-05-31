//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestResourceAcquireFromAssign — close 필요한 리소스 획득 AssignStmt 의 var 이름 반환 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestResourceAcquireFromAssign(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantVar  string
		wantCall bool
	}{
		{"single value open", "f := os.Open(\"x\")", "f", true},
		{"two value query", "rows, err := db.Query(\"q\")", "rows", true},
		{"blank discard", "_ := os.Open(\"x\")", "", false},
		{"non resource call", "v := compute()", "", false},
		{"not a call", "x := 1", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			name, call := resourceAcquireFromAssign(as)
			if name != tt.wantVar {
				t.Errorf("var = %q, want %q", name, tt.wantVar)
			}
			if (call != nil) != tt.wantCall {
				t.Errorf("call present=%v, want %v", call != nil, tt.wantCall)
			}
		})
	}
	if name, call := resourceAcquireFromAssign(nil); name != "" || call != nil {
		t.Fatal("nil assign should return empty")
	}
}
