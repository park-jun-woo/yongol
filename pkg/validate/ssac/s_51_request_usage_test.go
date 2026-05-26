//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-51 — OpenAPI request field 가 SSaC 에서 사용되지 않으면 경고

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS51RequestUsage(t *testing.T) {
	oapi := map[string]rule.StringSet{"OpenAPI.request.Create": {"Name": true, "Price": true}}
	t.Run("Fires_unused", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: "Name"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Passes_all_used_args", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Args: []ssac.Arg{{Source: "request", Field: "Name"}, {Source: "request", Field: "Price"}}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Passes_via_inputs", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Inputs: map[string]string{"n": "request.Name", "p": "request.Price"}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertNoDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Passes_via_verify_password", func(t *testing.T) {
		lk := map[string]rule.StringSet{"OpenAPI.request.Login": {"Email": true, "Password": true}}
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Login", FileName: "a.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "verify-password", Line: 3, EmailExpr: "request.Email", PasswordExpr: "request.Password"},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: lk})
		assertNoDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Fires_input_non_request_prefix", func(t *testing.T) {
		// Input value "owner.ID" is not a request field — "Price" remains unused.
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Inputs: map[string]string{"n": "request.Name", "o": "owner.ID"}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Fires_input_request_dot_only", func(t *testing.T) {
		// Input value "request." has empty field after prefix — "Price" remains unused.
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "post", Line: 3, Model: "O.C", Inputs: map[string]string{"n": "request.Name", "x": "request."}},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: oapi})
		assertDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Fires_verify_password_request_dot_only", func(t *testing.T) {
		// EmailExpr is "request." (empty field) — "Email" remains unused.
		lk := map[string]rule.StringSet{"OpenAPI.request.Login": {"Email": true, "Password": true}}
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Login", FileName: "a.ssac", Line: 1, Sequences: []ssac.Sequence{
			{Type: "verify-password", Line: 3, EmailExpr: "request.", PasswordExpr: "request.Password"},
		}}}}
		fs.SetGround(&rule.Ground{Lookup: lk})
		assertDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Skips_subscribe", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "W", FileName: "w.ssac", Line: 1,
			Subscribe: &ssac.SubscribeInfo{Topic: "t", MessageType: "M"}}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{"OpenAPI.request.W": {"ID": true}}})
		assertNoDiag(t, s51RequestUsage(fs), "[S-51]")
	})
	t.Run("Skips_nil_ground", func(t *testing.T) {
		assertNoDiag(t, s51RequestUsage(&yongol.Fullstack{}), "[S-51]")
	})
	t.Run("Skips_no_oapi", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []ssac.ServiceFunc{{Name: "Create", FileName: "o.ssac", Line: 1}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})
		assertNoDiag(t, s51RequestUsage(fs), "[S-51]")
	})
}
