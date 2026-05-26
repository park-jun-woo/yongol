//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-46 — Result type name 대문자 시작 필수 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS46UppercaseStart(t *testing.T) {
	mk := func(typ string) *yongol.Fullstack {
		return &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Result: &ssac.Result{Type: typ, Var: "x"}},
		}}}}
	}
	t.Run("Fires_lowercase", func(t *testing.T) { assertDiag(t, s46UppercaseStart(mk("order")), "[S-46]") })
	t.Run("Passes_uppercase", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(mk("Order")), "[S-46]") })
	t.Run("Passes_primitive", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(mk("int64")), "[S-46]") })
	t.Run("Passes_slice", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(mk("[]Order")), "[S-46]") })
	t.Run("Skips_nil_result", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "put", Line: 3, Result: nil},
		}}}}
		assertNoDiag(t, s46UppercaseStart(fs), "[S-46]")
	})
	t.Run("Skips_empty_type", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(mk("")), "[S-46]") })
	t.Run("Skips_bare_slice", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(mk("[]")), "[S-46]") })
	t.Run("Empty", func(t *testing.T) { assertNoDiag(t, s46UppercaseStart(&yongol.Fullstack{}), "[S-46]") })
}
