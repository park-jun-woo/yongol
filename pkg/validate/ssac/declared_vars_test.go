//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what declaredVars — 선언된 result 변수 집합 검증 (subscribe/normal/empty/upto)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestDeclaredVars(t *testing.T) {
	t.Run("empty function", func(t *testing.T) {
		fn := parsessac.ServiceFunc{}
		vars := declaredVars(fn, 0)
		if len(vars) != 0 {
			t.Fatalf("expected 0, got %v", vars)
		}
	})

	t.Run("subscribe adds message", func(t *testing.T) {
		fn := parsessac.ServiceFunc{
			Subscribe: &parsessac.SubscribeInfo{},
		}
		vars := declaredVars(fn, 0)
		if !vars["message"] {
			t.Error("expected message in vars")
		}
	})

	t.Run("collects result vars up to index", func(t *testing.T) {
		fn := parsessac.ServiceFunc{
			Sequences: []parsessac.Sequence{
				{Result: &parsessac.Result{Var: "user"}},
				{Result: &parsessac.Result{Var: "order"}},
				{Result: &parsessac.Result{Var: "item"}},
			},
		}
		vars := declaredVars(fn, 2)
		if !vars["user"] {
			t.Error("expected user")
		}
		if !vars["order"] {
			t.Error("expected order")
		}
		if vars["item"] {
			t.Error("item should not be included (upto=2)")
		}
	})

	t.Run("nil result skipped", func(t *testing.T) {
		fn := parsessac.ServiceFunc{
			Sequences: []parsessac.Sequence{
				{Result: nil},
				{Result: &parsessac.Result{Var: "user"}},
			},
		}
		vars := declaredVars(fn, 2)
		if len(vars) != 1 {
			t.Fatalf("expected 1, got %v", vars)
		}
		if !vars["user"] {
			t.Error("expected user")
		}
	})

	t.Run("empty var skipped", func(t *testing.T) {
		fn := parsessac.ServiceFunc{
			Sequences: []parsessac.Sequence{
				{Result: &parsessac.Result{Var: ""}},
			},
		}
		vars := declaredVars(fn, 1)
		if len(vars) != 0 {
			t.Fatalf("expected 0, got %v", vars)
		}
	})
}
