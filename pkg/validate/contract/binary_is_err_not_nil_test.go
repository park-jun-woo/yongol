//ff:func feature=validate-contract type=test control=iteration dimension=1 topic=preserve-safety
//ff:what TestBinaryIsErrNotNil — `x != nil` / `nil != x` err 패턴 판정 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestBinaryIsErrNotNil(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		errName string
		want    bool
	}{
		{"err lhs any", "err != nil", "", true},
		{"err rhs any", "nil != err", "", true},
		{"suffix err any", "dbErr != nil", "", true},
		{"non err any", "x != nil", "", false},
		{"pinned match", "dbErr != nil", "dbErr", true},
		{"pinned mismatch", "err != nil", "dbErr", false},
		{"selector excluded", "foo.err != nil", "", false},
		{"both nil", "nil != nil", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bin := mustExpr(t, tt.src).(*ast.BinaryExpr)
			if got := binaryIsErrNotNil(bin, tt.errName); got != tt.want {
				t.Fatalf("binaryIsErrNotNil(%q, %q) = %v, want %v", tt.src, tt.errName, got, tt.want)
			}
		})
	}
}
