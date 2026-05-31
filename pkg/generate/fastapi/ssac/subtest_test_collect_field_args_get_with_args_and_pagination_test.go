//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestCollectFieldArgsGetWithArgsAndPagination — GetWithArgsAndPagination 서브테스트
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestCollectFieldArgsGetWithArgsAndPagination(t *testing.T) {

	op := ir.Op{Kind: ir.OpGet, Get: &ir.GetOp{
		Args:           []ir.FieldArg{faKey("a")},
		PaginationArgs: []ir.FieldArg{faKey("cursor")},
	}}
	got := collectFieldArgs(op)
	if len(got) != 2 || got[0].Key != "a" || got[1].Key != "cursor" {
		t.Errorf("got %v, want [a cursor]", got)
	}

}
