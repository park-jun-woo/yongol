package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_GetOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpGet,
		Get: &ir.GetOp{
			Args:           []ir.FieldArg{{Key: "id", Location: ir.LocPath}},
			PaginationArgs: []ir.FieldArg{{Key: "cursor", Location: ir.LocQuery}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].Key != "id" {
		t.Errorf("got[0].Key = %q, want %q", got[0].Key, "id")
	}
	if got[1].Key != "cursor" {
		t.Errorf("got[1].Key = %q, want %q", got[1].Key, "cursor")
	}
}

func TestCollectFieldArgs_PostOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPost,
		Post: &ir.PostOp{
			Args: []ir.FieldArg{{Key: "title", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].Key != "title" {
		t.Errorf("got[0].Key = %q, want %q", got[0].Key, "title")
	}
}

func TestCollectFieldArgs_PutOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPut,
		Put: &ir.PutOp{
			Args: []ir.FieldArg{{Key: "name", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_DeleteOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpDelete,
		Delete: &ir.DeleteOp{
			Args: []ir.FieldArg{{Key: "id", Location: ir.LocPath}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_AuthOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpAuth,
		Auth: &ir.AuthOp{
			Inputs: []ir.FieldArg{{Key: "org_id", Location: ir.LocPath}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_StateOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpState,
		State: &ir.StateOp{
			Inputs: []ir.FieldArg{{Key: "status", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_CallOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpCall,
		Call: &ir.CallOp{
			Args: []ir.FieldArg{{Key: "input", Location: ir.LocVar}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_EvalOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpEval,
		Eval: &ir.EvalOp{
			Args: []ir.FieldArg{{Key: "x", Location: ir.LocLiteral}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}

func TestCollectFieldArgs_PublishOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPublish,
		Publish: &ir.PublishOp{
			Payload: []ir.FieldArg{{Key: "event_type"}},
			Options: []ir.FieldArg{{Key: "delay"}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
}

func TestCollectFieldArgs_EmptyOp(t *testing.T) {
	op := ir.Op{
		Kind:  ir.OpEmpty,
		Empty: &ir.EmptyOp{VarName: "item"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("EmptyOp should return nil, got %v", got)
	}
}

func TestCollectFieldArgs_ExistsOp(t *testing.T) {
	op := ir.Op{
		Kind:   ir.OpExists,
		Exists: &ir.ExistsOp{VarName: "existing"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("ExistsOp should return nil, got %v", got)
	}
}

func TestCollectFieldArgs_ResponseOp(t *testing.T) {
	op := ir.Op{
		Kind:     ir.OpResponse,
		Response: &ir.ResponseOp{SingleVar: "result"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("ResponseOp should return nil, got %v", got)
	}
}

func TestCollectFieldArgs_NilPointer(t *testing.T) {
	// Nil inner pointer should not panic
	op := ir.Op{Kind: ir.OpGet, Get: nil}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("nil Get pointer should return nil, got %v", got)
	}
}
