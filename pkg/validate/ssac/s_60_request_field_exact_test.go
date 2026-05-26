//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-60 — request.<field> 가 OpenAPI 스키마에 정확히 존재해야 함

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS60RequestFieldExact(t *testing.T) {
	oapi := map[string]rule.StringSet{"OpenAPI.request.Create": {"Name": true, "Price": true}}
	t.Run("Fires_unknown_field_args", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "o.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "Bogus"}}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Fires_unknown_field_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "o.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Inputs: map[string]string{"n": "request.Bogus"}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Passes_valid_field", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "o.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "Name"}}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Passes_param_field", func(t *testing.T) {
		lk := map[string]rule.StringSet{"OpenAPI.param.GetOne": {"id": true}}
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "GetOne", FileName: "g.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "get", Line: 3, Args: []ssac.Arg{{Source: "request", Field: "id"}}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: lk})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Skips_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "W", FileName: "w.ssac", Line: 1, Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s60RequestFieldExact(&yongol.Fullstack{}), "[S-60]")
	})
	t.Run("Skips_no_expected", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "X", FileName: "x.ssac", Line: 1}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Skips_input_non_request", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "o.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Inputs: map[string]string{"n": "owner.ID"}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
	t.Run("Skips_request_dot_only", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Create", FileName: "o.ssac", Line: 1,
			Sequences: []ssac.Sequence{{Type: "post", Line: 3, Inputs: map[string]string{"n": "request."}}},
		}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s60RequestFieldExact(fs), "[S-60]")
	})
}
