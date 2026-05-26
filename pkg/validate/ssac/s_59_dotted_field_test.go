//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-59 — variable.field 참조 시 필드가 변수 타입에 실재해야 함

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS59DottedField(t *testing.T) {
	schema := map[string][]string{"SSaC.var.Cancel.res": {"ID", "Status"}}
	t.Run("Fires_unknown_field_in_args", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Cancel", FileName: "r.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "pkg.Func",
				Args: []ssac.Arg{{Source: "res", Field: "Bogus"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: schema, Types: map[string]string{}})
		assertDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Fires_unknown_field_in_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Cancel", FileName: "r.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "state", Line: 3, Inputs: map[string]string{"status": "res.Bogus"}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: schema, Types: map[string]string{}})
		assertDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Passes_valid_field_args", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Cancel", FileName: "r.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "pkg.Func",
				Args: []ssac.Arg{{Source: "res", Field: "Status"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: schema, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Passes_valid_field_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Cancel", FileName: "r.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "state", Line: 3, Inputs: map[string]string{"status": "res.Status"}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: schema, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Skips_reserved_source", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "Name"}}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Skips_unregistered_var", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, Model: "pkg.Func", Args: []ssac.Arg{{Source: "unknown", Field: "F"}}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s59DottedField(&yongol.Fullstack{}), "[S-59]")
	})
	t.Run("Skips_empty_source_or_field", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, Args: []ssac.Arg{{Source: "", Field: "F"}, {Source: "v", Field: ""}}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{"SSaC.var.X.v": {"F"}}, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Skips_unregistered_var_in_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Inputs: map[string]string{"v": "unknown.Field"}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
	t.Run("Skips_reserved_source_in_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Inputs: map[string]string{"v": "query.Cursor"}}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})
}
