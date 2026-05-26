//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-49 — Model.Method 의 Method 가 symbol table 에 존재해야 한다

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS49SymbolTableMethod(t *testing.T) {
	mkFS := func(model string) *yongol.Fullstack {
		return &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: model},
		}}}}
	}
	methods := map[string]rule.StringSet{"SymbolTable.method.Order": {"FindByID": true}}
	t.Run("Fires_unknown", func(t *testing.T) {
		fs := mkFS("Order.Unknown")
		fs.SetGround(&rule.Ground{Lookup: methods})
		assertDiag(t, s49SymbolTableMethod(fs), "[S-49]")
	})
	t.Run("Passes_known", func(t *testing.T) {
		fs := mkFS("Order.FindByID")
		fs.SetGround(&rule.Ground{Lookup: methods})
		assertNoDiag(t, s49SymbolTableMethod(fs), "[S-49]")
	})
	t.Run("Skips_no_set", func(t *testing.T) {
		fs := mkFS("Order.FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s49SymbolTableMethod(fs), "[S-49]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s49SymbolTableMethod(mkFS("Order.FindByID")), "[S-49]")
	})
	t.Run("Skips_non_crud", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "empty", Line: 3, Target: "order"},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: methods})
		assertNoDiag(t, s49SymbolTableMethod(fs), "[S-49]")
	})
	t.Run("Skips_no_dot", func(t *testing.T) {
		fs := mkFS("FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s49SymbolTableMethod(fs), "[S-49]")
	})
	t.Run("Empty", func(t *testing.T) { assertNoDiag(t, s49SymbolTableMethod(&yongol.Fullstack{}), "[S-49]") })
}
