//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-71 — SSaC Input 변수 scope 밖 참조 시 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS71UnknownVariable(t *testing.T) {
	t.Run("Fires_unknown_var", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "put", Line: 3,
				Inputs: map[string]string{"id": "bogus.ID"},
			}},
		}}}
		assertDiag(t, s71UnknownVariable(fs), "[S-71]")
	})
	t.Run("Passes_request_ref", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{
				Type: "post", Line: 3,
				Inputs: map[string]string{"name": "request.Name"},
			}},
		}}}
		assertNoDiag(t, s71UnknownVariable(fs), "[S-71]")
	})
	t.Run("Passes_result_var_ref", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Result: &ssac.Result{Var: "course", Type: "Course"}},
				{Type: "put", Line: 5, Inputs: map[string]string{"id": "course.ID"}},
			},
		}}}
		assertNoDiag(t, s71UnknownVariable(fs), "[S-71]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s71UnknownVariable(&yongol.Fullstack{}), "[S-71]")
	})
}
