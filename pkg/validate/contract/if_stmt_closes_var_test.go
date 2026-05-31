//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIfStmtClosesVar — IfStmt body 내에 varName Close 호출 존재 여부 검증
package contract

import (
	"testing"
)

func TestIfStmtClosesVar(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		varName string
		want    bool
	}{
		{"close in if", "if f != nil { f.Close() }", "f", true},
		{"wrong var in if", "if f != nil { f.Close() }", "g", false},
		{"no close in if", "if f != nil { f.Read(b) }", "f", false},
		{"not an if", "f.Close()", "f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := mustFirstStmt(t, tt.body)
			if got := ifStmtClosesVar(stmt, tt.varName); got != tt.want {
				t.Fatalf("ifStmtClosesVar(%q, %q) = %v, want %v", tt.body, tt.varName, got, tt.want)
			}
		})
	}
}
