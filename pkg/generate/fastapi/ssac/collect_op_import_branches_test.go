//ff:func feature=gen-fastapi type=test control=selection
//ff:what TestCollectOpImportBranches — Get/Post/Put/Delete/Publish/VerifyPW/Call/Eval 분기 커버

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func newImportData() importData {
	return importData{
		Models:  make(map[string]bool),
		ExtPkgs: make(map[string]map[string]bool),
	}
}

func TestCollectOpImportBranches(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpGet, Get: &ir.GetOp{Model: "User"}}, "f")
		if !d.UsesSelect || !d.Models["User"] {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("GetNil", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpGet}, "f")
		if !d.UsesSelect || len(d.Models) != 0 {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("Post", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpPost, Post: &ir.PostOp{Model: "Item"}}, "f")
		if !d.Models["Item"] {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("PostNil", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpPost}, "f")
		if len(d.Models) != 0 {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("Put", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpPut, Put: &ir.PutOp{Model: "Item"}}, "f")
		if !d.UsesUpdate || !d.Models["Item"] {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("PutNil", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpPut}, "f")
		if !d.UsesUpdate || len(d.Models) != 0 {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "Item"}}, "f")
		if !d.UsesDelete || !d.Models["Item"] {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("DeleteNil", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpDelete}, "f")
		if !d.UsesDelete || len(d.Models) != 0 {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("Publish", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpPublish}, "f")
		if !d.HasPublish {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("VerifyPassword", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{Model: "User"}}, "f")
		if !d.UsesSelect || !d.Models["User"] {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("VerifyPasswordNil", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpVerifyPassword}, "f")
		if !d.UsesSelect || len(d.Models) != 0 {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("CallExternal", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "other", Function: "Do"}}, "f")
		if !d.ExtPkgs["other"]["Do"] {
			t.Errorf("got %+v", d.ExtPkgs)
		}
	})
	t.Run("CallSelfSkipped", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "f", Function: "Do"}}, "f")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("expected self-import skipped, got %+v", d.ExtPkgs)
		}
	})
	t.Run("CallNilOrEmptyPackage", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpCall}, "f")
		collectOpImport(&d, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: ""}}, "f")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("got %+v", d.ExtPkgs)
		}
	})
	t.Run("EvalExternal", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "other", Function: "Calc"}}, "f")
		if !d.ExtPkgs["other"]["Calc"] {
			t.Errorf("got %+v", d.ExtPkgs)
		}
	})
	t.Run("EvalSelfSkipped", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "f", Function: "Calc"}}, "f")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("got %+v", d.ExtPkgs)
		}
	})
	t.Run("EvalNilOrEmpty", func(t *testing.T) {
		d := newImportData()
		collectOpImport(&d, ir.Op{Kind: ir.OpEval}, "f")
		if len(d.ExtPkgs) != 0 {
			t.Errorf("got %+v", d.ExtPkgs)
		}
	})
}
