//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-48 — CRUD Model 이 symbol table 에 존재해야 한다

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS48SymbolTableModel(t *testing.T) {
	mkFS := func(model string) *yongol.Fullstack {
		return &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: model},
		}}}}
	}
	t.Run("Fires_unknown", func(t *testing.T) {
		fs := mkFS("Unknown.FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"SymbolTable.model": {"Order": true}}})
		assertDiag(t, s48SymbolTableModel(fs), "[S-48]")
	})
	t.Run("Passes_known", func(t *testing.T) {
		fs := mkFS("Order.FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"SymbolTable.model": {"Order": true}}})
		assertNoDiag(t, s48SymbolTableModel(fs), "[S-48]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s48SymbolTableModel(mkFS("Order.FindByID")), "[S-48]")
	})
	t.Run("Skips_no_table", func(t *testing.T) {
		fs := mkFS("Order.FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s48SymbolTableModel(fs), "[S-48]")
	})
	t.Run("Skips_non_crud", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "empty", Line: 3, Target: "order"},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"SymbolTable.model": {}}})
		assertNoDiag(t, s48SymbolTableModel(fs), "[S-48]")
	})
	t.Run("Skips_no_dot", func(t *testing.T) {
		fs := mkFS("FindByID")
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"SymbolTable.model": {}}})
		assertNoDiag(t, s48SymbolTableModel(fs), "[S-48]")
	})
	t.Run("Empty", func(t *testing.T) { assertNoDiag(t, s48SymbolTableModel(&yongol.Fullstack{}), "[S-48]") })
}
