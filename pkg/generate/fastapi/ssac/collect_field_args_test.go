//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectFieldArgs — Op 종류별 FieldArg 추출 (nil/모든 Kind 분기 커버)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs(t *testing.T) {
	fa := func(k string) ir.FieldArg { return ir.FieldArg{Key: k} }

	t.Run("GetWithArgsAndPagination", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpGet, Get: &ir.GetOp{
			Args:           []ir.FieldArg{fa("a")},
			PaginationArgs: []ir.FieldArg{fa("cursor")},
		}}
		got := collectFieldArgs(op)
		if len(got) != 2 || got[0].Key != "a" || got[1].Key != "cursor" {
			t.Errorf("got %v, want [a cursor]", got)
		}
	})
	t.Run("GetNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpGet}) != nil {
			t.Error("expected nil for nil Get")
		}
	})
	t.Run("Post", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpPost, Post: &ir.PostOp{Args: []ir.FieldArg{fa("p")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "p" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("PostNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpPost}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("Put", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpPut, Put: &ir.PutOp{Args: []ir.FieldArg{fa("u")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "u" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("PutNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpPut}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("Delete", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Args: []ir.FieldArg{fa("d")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "d" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("DeleteNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpDelete}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("Auth", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpAuth, Auth: &ir.AuthOp{Inputs: []ir.FieldArg{fa("au")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "au" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("AuthNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpAuth}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("State", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpState, State: &ir.StateOp{Inputs: []ir.FieldArg{fa("s")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "s" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("StateNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpState}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("Call", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Args: []ir.FieldArg{fa("c")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "c" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("CallNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpCall}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("Eval", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Args: []ir.FieldArg{fa("e")}}}
		if got := collectFieldArgs(op); len(got) != 1 || got[0].Key != "e" {
			t.Errorf("got %v", got)
		}
	})
	t.Run("EvalNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpEval}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("PublishPayloadAndOptions", func(t *testing.T) {
		op := ir.Op{Kind: ir.OpPublish, Publish: &ir.PublishOp{
			Payload: []ir.FieldArg{fa("pl")},
			Options: []ir.FieldArg{fa("op")},
		}}
		got := collectFieldArgs(op)
		if len(got) != 2 || got[0].Key != "pl" || got[1].Key != "op" {
			t.Errorf("got %v, want [pl op]", got)
		}
	})
	t.Run("PublishNil", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpPublish}) != nil {
			t.Error("expected nil")
		}
	})
	t.Run("DefaultUnknownKind", func(t *testing.T) {
		if collectFieldArgs(ir.Op{Kind: ir.OpKind(9999)}) != nil {
			t.Error("expected nil for unknown kind")
		}
	})
}
