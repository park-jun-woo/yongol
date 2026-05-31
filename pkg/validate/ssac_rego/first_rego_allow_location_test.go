//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestFirstRegoAllowLocation(t *testing.T) {
	policies := []rego.Policy{{
		File:  "p.rego",
		Rules: []rego.AllowRule{{Actions: []string{"read"}, Resource: "x", SourceLine: 7}},
	}}
	locs := firstRegoAllowLocation(policies)
	if locs[[2]string{"read", "x"}].Line != 7 {
		t.Errorf("locs = %+v", locs)
	}
}
