//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateRego/populateRegoPolicy — allow rule auth 쌍, roles, claims 등록

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// TestPopulateRego_AuthPairsRolesClaims covers the primary policy fields.
func TestPopulateRego_AuthPairsRolesClaims(t *testing.T) {
	pol := rego.Policy{
		File: "policies/project.rego",
		Rules: []rego.AllowRule{
			{Actions: []string{"read", "delete"}, Resource: "project", UsesRole: true, RoleValue: "admin"},
		},
		ClaimsRefs: []string{"org_id", "user_id"},
	}
	fs := newMinimalFullstack(withParsedPolicies(pol))
	g := newGround()

	populateRego(g, fs)

	pairs := g.Pairs["Policy.auth"]
	if !pairs["read:project"] || !pairs["delete:project"] {
		t.Errorf("Policy.auth missing read/delete:project: %v", pairs)
	}
	roles := g.Lookup["Rego.roles"]
	if !roles["admin"] {
		t.Errorf("Rego.roles missing admin: %v", roles)
	}
	claims := g.Lookup["Rego.claims"]
	if !claims["org_id"] || !claims["user_id"] {
		t.Errorf("Rego.claims missing: %v", claims)
	}
}
