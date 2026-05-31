//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestExprStmtClosesVar — ExprStmt/AssignStmt 가 varName.Close() 호출인지 검증
package contract

import (
	"testing"
)

func TestExprStmtClosesVar(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		varName string
		want    bool
	}{
		{"bare close", "f.Close()", "f", true},
		{"assign close", "err = f.Close()", "f", true},
		{"wrong var", "f.Close()", "g", false},
		{"non close expr", "f.Read(b)", "f", false},
		{"non expr stmt", "return", "f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stmt := mustFirstStmt(t, tt.body)
			if got := exprStmtClosesVar(stmt, tt.varName); got != tt.want {
				t.Fatalf("exprStmtClosesVar(%q, %q) = %v, want %v", tt.body, tt.varName, got, tt.want)
			}
		})
	}
}
