//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what s36TrackAssignment — Result.Var→Type 기록 + stale 초기화 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS36TrackAssignment(t *testing.T) {
	t.Run("Records", func(t *testing.T) {
		seq := parsessac.Sequence{
			Type:   "get",
			Result: &parsessac.Result{Type: "Order", Var: "order"},
		}
		varType := map[string]string{}
		stale := map[string]bool{"order": true}
		s36TrackAssignment(seq, varType, stale)
		if varType["order"] != "Order" {
			t.Errorf("varType[order] = %q, want Order", varType["order"])
		}
		if stale["order"] {
			t.Error("stale[order] should be reset to false after assignment")
		}
	})
	t.Run("NilResultSkipped", func(t *testing.T) {
		seq := parsessac.Sequence{Type: "get", Result: nil}
		varType := map[string]string{}
		stale := map[string]bool{}
		s36TrackAssignment(seq, varType, stale)
		if len(varType) != 0 {
			t.Error("expected no varType entries for nil Result")
		}
	})
	t.Run("EmptyVarSkipped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Type:   "get",
			Result: &parsessac.Result{Type: "Order", Var: ""},
		}
		varType := map[string]string{}
		stale := map[string]bool{}
		s36TrackAssignment(seq, varType, stale)
		if len(varType) != 0 {
			t.Error("expected no varType entries for empty Var")
		}
	})
	t.Run("SliceTypeStripped", func(t *testing.T) {
		seq := parsessac.Sequence{
			Type:   "get",
			Result: &parsessac.Result{Type: "[]Order", Var: "orders"},
		}
		varType := map[string]string{}
		stale := map[string]bool{}
		s36TrackAssignment(seq, varType, stale)
		if varType["orders"] != "Order" {
			t.Errorf("varType[orders] = %q, want Order ([] prefix stripped)", varType["orders"])
		}
	})
}
