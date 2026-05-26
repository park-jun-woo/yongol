//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-58 — ErrStatus 가 유효한 HTTP 상태 코드(100-599)가 아니면 에러

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS58InvalidErrStatus(t *testing.T) {
	t.Run("Fires_status_below_100", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "empty", Line: 3, ErrStatus: 99}},
		}}}
		assertDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Fires_status_above_599", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, ErrStatus: 600}},
		}}}
		assertDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Passes_valid_4xx", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "auth", Line: 3, ErrStatus: 403}},
		}}}
		assertNoDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Passes_valid_5xx", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, ErrStatus: 500}},
		}}}
		assertNoDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Skips_zero_status", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "empty", Line: 3, ErrStatus: 0}},
		}}}
		assertNoDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Passes_boundary_100", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, ErrStatus: 100}},
		}}}
		assertNoDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Passes_boundary_599", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "X", FileName: "x.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "call", Line: 3, ErrStatus: 599}},
		}}}
		assertNoDiag(t, s58InvalidErrStatus(fs), "[S-58]")
	})
	t.Run("Empty_funcs", func(t *testing.T) {
		assertNoDiag(t, s58InvalidErrStatus(&yongol.Fullstack{}), "[S-58]")
	})
}
