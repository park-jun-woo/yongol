//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestUnmarshalInStmt — 미가드 Unmarshal DF-01 + if-init/assign-guard/무매칭 분기 + 다중 target
package qcheck

import (
	"go/token"
	"testing"
)

func TestUnmarshalInStmt(t *testing.T) {
	fset := token.NewFileSet()
	targets := []string{"json"}

	t.Run("UnguardedFinding", func(t *testing.T) {
		list := blockStmts(t, "_ = json.Unmarshal(b, v)")
		got := unmarshalInStmt(list[0], list, 0, targets, fset)
		if len(got) != 1 || got[0].Category != "DF-01" || got[0].Detail != "json.Unmarshal" {
			t.Fatalf("expected one DF-01 finding, got %+v", got)
		}
	})

	t.Run("IfInitGuard", func(t *testing.T) {
		list := blockStmts(t, "if err := json.Unmarshal(b, v); err != nil { return }")
		if got := unmarshalInStmt(list[0], list, 0, targets, fset); len(got) != 0 {
			t.Errorf("expected no findings for if-init guard, got %+v", got)
		}
	})

	t.Run("AssignGuard", func(t *testing.T) {
		list := blockStmts(t, "err := json.Unmarshal(b, v)\nif err != nil { return }")
		if got := unmarshalInStmt(list[0], list, 0, targets, fset); len(got) != 0 {
			t.Errorf("expected no findings for assign guard, got %+v", got)
		}
	})

	t.Run("NoCall", func(t *testing.T) {
		list := blockStmts(t, "x := 1")
		if got := unmarshalInStmt(list[0], list, 0, targets, fset); len(got) != 0 {
			t.Errorf("expected no findings when no call present, got %+v", got)
		}
	})

	t.Run("MultiTargets", func(t *testing.T) {
		// yaml target present, json absent in stmt -> one finding for yaml only.
		list := blockStmts(t, "_ = yaml.Unmarshal(b, v)")
		got := unmarshalInStmt(list[0], list, 0, []string{"json", "yaml"}, fset)
		if len(got) != 1 || got[0].Detail != "yaml.Unmarshal" {
			t.Fatalf("expected one yaml.Unmarshal finding, got %+v", got)
		}
	})
}
