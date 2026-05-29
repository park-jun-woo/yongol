//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xnp63RoleManifest — nil/no ground/declared role/undeclared role 검증

package rego_manifest

import (
	"strings"
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnp63RoleManifest_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := xnp63RoleManifest(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xnp63RoleManifest(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("declared role passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true, "user": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{{UsesRole: true, RoleValue: "admin"}}},
			},
		}
		fs.SetGround(g)
		diags := xnp63RoleManifest(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("undeclared role fires XNP-63", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{{UsesRole: true, RoleValue: "superadmin"}}},
			},
		}
		fs.SetGround(g)
		diags := xnp63RoleManifest(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XNP-63]") {
			t.Errorf("expected XNP-63, got %s", diags[0].Message)
		}
	})
}

func TestXnp63RoleManifest(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
