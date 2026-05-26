//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-57 — @call input type 불일치 시 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS57FuncRequestType(t *testing.T) {
	t.Run("Fires_type_mismatch", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "AcceptGig", FileName: "gig.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "billing.HoldEscrow",
				Args: []ssac.Arg{{Source: "gig", Field: "Price"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{
			Schemas: map[string][]string{"Func.request.HoldEscrow": {"Price"}},
			Types: map[string]string{
				"SSaC.var.AcceptGig.gig":        "Gig",
				"Func.request.HoldEscrow.Price": "int64",
			},
		})
		assertDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Passes_type_match", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "AcceptGig", FileName: "gig.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "billing.HoldEscrow",
				Args: []ssac.Arg{{Source: "gig", Field: "Price"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{
			Schemas: map[string][]string{"Func.request.HoldEscrow": {"Price"}},
			Types: map[string]string{
				"SSaC.var.AcceptGig.gig":        "Gig",
				"Func.request.HoldEscrow.Price": "Gig",
			},
		})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Skips_non_call", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "GetCourse", FileName: "c.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "get", Line: 3, Model: "Course.FindByID"}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s57FuncRequestType(&yongol.Fullstack{}), "[S-57]")
	})
	t.Run("Skips_no_schema", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "DoStuff", FileName: "s.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "pkg.Func",
				Args: []ssac.Arg{{Source: "x", Field: "Y"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Skips_empty_source_or_field", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "pkg.Func",
				Args: []ssac.Arg{{Source: "", Field: "Y"}, {Source: "x", Field: ""}},
			}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{"Func.request.Func": {"Y"}}, Types: map[string]string{}})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Skips_unknown_types", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "pkg.Func",
				Args: []ssac.Arg{{Source: "y", Field: "Z"}},
			}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{"Func.request.Func": {"Z"}}, Types: map[string]string{}})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
	t.Run("Skips_model_without_dot", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, Model: "Func"}},
		}}}
		fs.SetGround(&rule.Ground{Schemas: map[string][]string{}, Types: map[string]string{}})
		assertNoDiag(t, s57FuncRequestType(fs), "[S-57]")
	})
}
