//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectRegoAllowPairs(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{{
			File: "policy.rego",
			Rules: []rego.AllowRule{
				{Actions: []string{"read", "delete"}, Resource: "project"},
			},
		}},
	}
	pairs := collectRegoAllowPairs(fs)
	if !pairs[[2]string{"read", "project"}] || !pairs[[2]string{"delete", "project"}] {
		t.Errorf("pairs = %v", pairs)
	}
	// nil fs → empty map, no panic.
	if got := collectRegoAllowPairs(nil); len(got) != 0 {
		t.Errorf("nil fs → %v", got)
	}
}
