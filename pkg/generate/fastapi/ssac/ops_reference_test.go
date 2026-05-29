//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestOpsReference — opsReferenceBody/Path/Query FieldArg Location 검사 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpsReferenceBody(t *testing.T) {
	t.Run("BodyArgPresent", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{
				Args: []ir.FieldArg{{Location: ir.LocBody, Key: "title"}},
			}},
		}
		if !opsReferenceBody(ops) {
			t.Error("expected true for LocBody arg")
		}
	})

	t.Run("NoBodyArg", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{
				Args: []ir.FieldArg{{Location: ir.LocPath, Key: "id"}},
			}},
		}
		if opsReferenceBody(ops) {
			t.Error("expected false when no LocBody arg")
		}
	})

	t.Run("NilOps", func(t *testing.T) {
		if opsReferenceBody(nil) {
			t.Error("expected false for nil ops")
		}
	})
}

func TestOpsReferencePath(t *testing.T) {
	t.Run("PathArgPresent", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{
				Args: []ir.FieldArg{{Location: ir.LocPath, Key: "id"}},
			}},
		}
		if !opsReferencePath(ops) {
			t.Error("expected true for LocPath arg")
		}
	})

	t.Run("NoPathArg", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{
				Args: []ir.FieldArg{{Location: ir.LocBody, Key: "title"}},
			}},
		}
		if opsReferencePath(ops) {
			t.Error("expected false when no LocPath arg")
		}
	})

	t.Run("NilOps", func(t *testing.T) {
		if opsReferencePath(nil) {
			t.Error("expected false for nil ops")
		}
	})
}

func TestOpsReferenceQuery(t *testing.T) {
	t.Run("QueryArgPresent", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{
				PaginationArgs: []ir.FieldArg{{Location: ir.LocQuery, Key: "per_page"}},
			}},
		}
		if !opsReferenceQuery(ops) {
			t.Error("expected true for LocQuery arg")
		}
	})

	t.Run("NoQueryArg", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{
				Args: []ir.FieldArg{{Location: ir.LocPath, Key: "id"}},
			}},
		}
		if opsReferenceQuery(ops) {
			t.Error("expected false when no LocQuery arg")
		}
	})

	t.Run("NilOps", func(t *testing.T) {
		if opsReferenceQuery(nil) {
			t.Error("expected false for nil ops")
		}
	})
}
