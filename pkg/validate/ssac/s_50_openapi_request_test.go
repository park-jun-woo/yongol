//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-50 — request.field 가 OpenAPI request schema 에 존재해야 한다

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS50OpenAPIRequest(t *testing.T) {
	oapi := map[string]rule.StringSet{"OpenAPI.request.CreateOrder": {"Name": true, "Price": true}}
	t.Run("Fires_missing", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "CreateOrder", FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: "Missing"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
	t.Run("Passes_known", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "CreateOrder", FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: "Name"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
	t.Run("Skips_non_request", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "CreateOrder", FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "get", Line: 3, Model: "O.F", Args: []ssac.Arg{{Source: "order", Field: "ID"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
	t.Run("Skips_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "W", FileName: "w.ssac",
			Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}, Sequences: []ssac.Sequence{
				{Type: "get", Line: 3, Model: "O.F", Args: []ssac.Arg{{Source: "request", Field: "ID"}}},
			}}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"OpenAPI.request.W": {}}})
		assertNoDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
	t.Run("Skips_no_oapi", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "CreateOrder", FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: "Name"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s50OpenAPIRequest(&yongol.Fullstack{}), "[S-50]")
	})
	t.Run("Skips_empty_field", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "CreateOrder", FileName: "o.ssac", Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: ""}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s50OpenAPIRequest(fs), "[S-50]")
	})
}
