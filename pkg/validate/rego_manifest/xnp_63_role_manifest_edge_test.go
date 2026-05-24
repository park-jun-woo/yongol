//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xnp63RoleManifestEdge — no roles in manifest/no UsesRole/duplicate role 검증

package rego_manifest

import (
	"strings"
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnp63RoleManifest_Edge(t *testing.T) {
	t.Run("manifest has no roles fires", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{{UsesRole: true, RoleValue: "admin"}}},
			},
		}
		fs.SetGround(g)
		diags := xnp63RoleManifest(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "not declared") {
			t.Errorf("expected undeclared message, got %s", diags[0].Message)
		}
	})

	t.Run("no UsesRole skips", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{{UsesRole: false}}},
			},
		}
		fs.SetGround(g)
		diags := xnp63RoleManifest(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("duplicate role deduplicates", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", Rules: []regoparser.AllowRule{
					{UsesRole: true, RoleValue: "unknown"},
					{UsesRole: true, RoleValue: "unknown"},
				}},
			},
		}
		fs.SetGround(g)
		diags := xnp63RoleManifest(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 (deduplicated), got %d: %+v", len(diags), diags)
		}
	})
}
