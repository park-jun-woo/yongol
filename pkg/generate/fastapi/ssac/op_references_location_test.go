//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestOpReferencesLocation — opReferencesLocation Op 의 FieldArg Location 참조 여부 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpReferencesLocation(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPost,
		Post: &ir.PostOp{Args: []ir.FieldArg{
			{Location: ir.LocBody},
			{Location: ir.LocPath},
		}},
	}

	if !opReferencesLocation(op, ir.LocBody) {
		t.Error("expected LocBody referenced")
	}
	if !opReferencesLocation(op, ir.LocPath) {
		t.Error("expected LocPath referenced")
	}
	if opReferencesLocation(op, ir.LocQuery) {
		t.Error("expected LocQuery not referenced")
	}

	empty := ir.Op{Kind: ir.OpPost, Post: nil}
	if opReferencesLocation(empty, ir.LocBody) {
		t.Error("expected false for op with no field args")
	}
}
