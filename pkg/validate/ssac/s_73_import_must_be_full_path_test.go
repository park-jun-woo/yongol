//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-73 — SSaC import 가 bare name 이면 에러 (full Go import path 필수)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS73ImportMustBeFullPath(t *testing.T) {
	t.Run("Fires_bare_name", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Imports: []string{"dashboard"},
		}}}
		assertDiag(t, s73ImportMustBeFullPath(fs), "[S-73]")
	})
	t.Run("Passes_full_path", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Imports: []string{"github.com/user/project/internal/billing"},
		}}}
		assertNoDiag(t, s73ImportMustBeFullPath(fs), "[S-73]")
	})
	t.Run("Empty_imports", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
		}}}
		assertNoDiag(t, s73ImportMustBeFullPath(fs), "[S-73]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s73ImportMustBeFullPath(&yongol.Fullstack{}), "[S-73]")
	})
}
