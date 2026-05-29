//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestIndexOfErrLhs — AssignStmt LHS 중 errName ident 위치 반환 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestIndexOfErrLhs(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		errName string
		want    int
	}{
		{"default err first", "err = f()", "", 0},
		{"default err second", "v, err := f()", "", 1},
		{"named match", "x, dbErr := f()", "dbErr", 1},
		{"no match", "x, y := f()", "", -1},
		{"named not present", "err = f()", "dbErr", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := mustFirstStmt(t, tt.body).(*ast.AssignStmt)
			if got := indexOfErrLhs(as, tt.errName); got != tt.want {
				t.Fatalf("indexOfErrLhs(%q, %q) = %d, want %d", tt.body, tt.errName, got, tt.want)
			}
		})
	}
}
