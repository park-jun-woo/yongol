//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIsErrCheckIf — IfStmt 이 err != nil 계열 가드인지 판정 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestIsErrCheckIf(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		errName string
		want    bool
	}{
		{"plain err guard", "if err != nil { return }", "", true},
		{"init err guard", "if err := f(); err != nil { return }", "", true},
		{"suffix err", "if dbErr != nil { return }", "", true},
		{"pinned match", "if dbErr != nil { return }", "dbErr", true},
		{"pinned mismatch", "if err != nil { return }", "dbErr", false},
		{"equality not neq", "if err == nil { return }", "", false},
		{"non err cond", "if x != nil { return }", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifs := mustFirstStmt(t, tt.body).(*ast.IfStmt)
			if got := isErrCheckIf(ifs, tt.errName); got != tt.want {
				t.Fatalf("isErrCheckIf(%q, %q) = %v, want %v", tt.body, tt.errName, got, tt.want)
			}
		})
	}
	if isErrCheckIf(nil, "") {
		t.Fatal("nil if stmt should return false")
	}
}
