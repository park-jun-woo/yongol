//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xpn54ClaimsToRego — manifest claim 참조 검증 (nil/no ground/pass/fire/middleware/openapi) 검증

package rego_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXpn54ClaimsToRego_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := xpn54ClaimsToRego(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no claims keys returns nil", func(t *testing.T) {
		g := &rule.Ground{Lookup: map[string]rule.StringSet{}}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("claim referenced in Rego passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"sub": true},
				"Rego.claims":          {"sub": true},
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("claim referenced in middleware passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"sub": true},
				"Middleware.claims":    {"sub": true},
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("claim referenced in OpenAPI response passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup:  map[string]rule.StringSet{"Manifest.claims.keys": {"sub": true}},
			Schemas: map[string][]string{
				"OpenAPI.response.GetUser": {"sub", "name"},
				"DDL.columns.users":        {"id", "name"}, // non-OpenAPI schema skipped
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("unreferenced claim fires XPN-54", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"orphan": true},
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn54ClaimsToRego(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XPN-54]") {
			t.Errorf("expected XPN-54, got %s", diags[0].Message)
		}
	})
}
