//ff:func feature=gen-ir type=test control=sequence
//ff:what TestRenameOpResultVar -- Op 종류별 result 변수명 충돌 해소 및 rename 등록 검증

package ir

import "testing"

func TestRenameOpResultVar(t *testing.T) {
	t.Run("get no collision keeps name", func(t *testing.T) {
		op := &Op{Kind: OpGet, Get: &GetOp{VarName: "user"}}
		declared := map[string]bool{}
		renames := map[string]string{}
		renameOpResultVar(op, declared, renames)
		if op.Get.VarName != "user" {
			t.Errorf("VarName = %q, want user", op.Get.VarName)
		}
		if len(renames) != 0 {
			t.Errorf("renames = %v, want empty", renames)
		}
	})

	t.Run("post collision renames and records", func(t *testing.T) {
		op := &Op{Kind: OpPost, Post: &PostOp{VarName: "user"}}
		declared := map[string]bool{"user": true}
		renames := map[string]string{}
		renameOpResultVar(op, declared, renames)
		if op.Post.VarName != "user_result" {
			t.Errorf("VarName = %q, want user_result", op.Post.VarName)
		}
		if renames["user"] != "user_result" {
			t.Errorf("renames[user] = %q, want user_result", renames["user"])
		}
	})

	t.Run("call result var", func(t *testing.T) {
		op := &Op{Kind: OpCall, Call: &CallOp{ResultVar: "r"}}
		declared := map[string]bool{"r": true}
		renames := map[string]string{}
		renameOpResultVar(op, declared, renames)
		if op.Call.ResultVar != "r_result" {
			t.Errorf("ResultVar = %q, want r_result", op.Call.ResultVar)
		}
	})

	t.Run("verify password result var", func(t *testing.T) {
		op := &Op{Kind: OpVerifyPassword, VerifyPW: &VerifyPasswordOp{ResultVar: "user"}}
		declared := map[string]bool{"user": true}
		renames := map[string]string{}
		renameOpResultVar(op, declared, renames)
		if op.VerifyPW.ResultVar != "user_result" {
			t.Errorf("ResultVar = %q, want user_result", op.VerifyPW.ResultVar)
		}
	})

	t.Run("other kind is a no-op", func(t *testing.T) {
		op := &Op{Kind: OpEmpty}
		renames := map[string]string{}
		renameOpResultVar(op, map[string]bool{}, renames)
		if len(renames) != 0 {
			t.Errorf("renames = %v, want empty for non-result kind", renames)
		}
	})
}
