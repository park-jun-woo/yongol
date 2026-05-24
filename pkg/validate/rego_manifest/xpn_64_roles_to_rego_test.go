//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xpn64RolesToRego — manifest role 참조 검증 (nil/no ground/pass/fire) 검증

package rego_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXpn64RolesToRego_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := xpn64RolesToRego(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xpn64RolesToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no manifest roles returns nil", func(t *testing.T) {
		g := &rule.Ground{Lookup: map[string]rule.StringSet{}}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn64RolesToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("all roles referenced passes", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true},
				"Rego.roles":     {"admin": true},
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn64RolesToRego(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("unused role fires XPN-64", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"Manifest.roles": {"admin": true, "superadmin": true},
				"Rego.roles":     {"admin": true},
			},
		}
		fs := &yongol.Fullstack{}
		fs.SetGround(g)
		diags := xpn64RolesToRego(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XPN-64]") {
			t.Errorf("expected XPN-64, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "superadmin") {
			t.Errorf("expected role name, got %s", diags[0].Message)
		}
	})
}
