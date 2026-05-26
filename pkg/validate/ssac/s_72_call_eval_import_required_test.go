//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-72 — @call/@eval 패키지 참조 시 import 선언 없으면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS72CallEvalImportRequired(t *testing.T) {
	t.Run("Fires_missing_import_call", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "billing.HoldEscrow",
			}},
		}}}
		assertDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Fires_missing_import_eval", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "eval", Line: 3, Model: "billing.IsZero",
			}},
		}}}
		assertDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Passes_with_import", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Imports: []string{"pkg/billing"},
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "billing.HoldEscrow",
			}},
		}}}
		assertNoDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Skips_non_call_eval", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "get", Line: 3, Model: "Course.FindByID",
			}},
		}}}
		assertNoDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Skips_empty_model", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "",
			}},
		}}}
		assertNoDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Skips_no_dot", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "call", Line: 3, Model: "Func",
			}},
		}}}
		assertNoDiag(t, s72CallEvalImportRequired(fs), "[S-72]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s72CallEvalImportRequired(&yongol.Fullstack{}), "[S-72]")
	})
}
