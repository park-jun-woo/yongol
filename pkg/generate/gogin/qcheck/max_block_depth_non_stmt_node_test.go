//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMaxBlockDepth — for/range/switch/typeswitch/select/case 각 노드 종류별 depth 검증
package qcheck

import (
	"go/ast"
	"testing"
)

func TestMaxBlockDepth_NonStmtNode(t *testing.T) {
	// A non-control node returns the input depth unchanged.
	if got := maxBlockDepth(&ast.Ident{Name: "x"}, 7); got != 7 {
		t.Errorf("maxBlockDepth(ident, 7) = %d, want 7", got)
	}
}
