//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpReferencesLocation — Op 의 FieldArg 중 지정 Location 참조 포함 여부 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpReferencesLocation(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpGet,
		Get: &ir.GetOp{
			Args:           []ir.FieldArg{{Key: "id", Location: ir.LocPath}},
			PaginationArgs: []ir.FieldArg{{Key: "cursor", Location: ir.LocQuery}},
		},
	}

	if !opReferencesLocation(op, ir.LocPath) {
		t.Error("expected true for LocPath")
	}
	if !opReferencesLocation(op, ir.LocQuery) {
		t.Error("expected true for LocQuery")
	}
	if opReferencesLocation(op, ir.LocBody) {
		t.Error("expected false for LocBody (not referenced)")
	}
}
