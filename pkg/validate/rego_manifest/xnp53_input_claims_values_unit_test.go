//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xnp53InputClaimsValues — claims ref 검증 (nil/no ground/pass/fire) 검증
package rego_manifest

import (
	"strings"
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnp53InputClaimsValues_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := xnp53InputClaimsValues(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xnp53InputClaimsValues(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no claims keys returns nil", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xnp53InputClaimsValues(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("declared ref passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"sub": true, "org_id": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", ClaimsRefs: []string{"sub", "org_id"}},
			},
		}
		fs.SetGround(g)
		diags := xnp53InputClaimsValues(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("duplicate undeclared ref deduplicates", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"sub": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", ClaimsRefs: []string{"unknown", "unknown"}},
			},
		}
		fs.SetGround(g)
		diags := xnp53InputClaimsValues(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 (deduplicated), got %d: %+v", len(diags), diags)
		}
	})

	t.Run("undeclared ref fires XNP-53", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.claims.keys": {"sub": true},
			},
		}
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego", ClaimsRefs: []string{"unknown_claim"}},
			},
		}
		fs.SetGround(g)
		diags := xnp53InputClaimsValues(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XNP-53]") {
			t.Errorf("expected XNP-53, got %s", diags[0].Message)
		}
	})
}
