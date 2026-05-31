//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestScanRegoAllowLocations(t *testing.T) {
	locs := map[[2]string]PairLocation{}
	p := rego.Policy{
		File: "p.rego",
		Rules: []rego.AllowRule{
			{Actions: []string{"read"}, Resource: "x", SourceLine: 10},
			{Actions: []string{"read"}, Resource: "x", SourceLine: 20}, // dup, first wins
		},
	}
	scanRegoAllowLocations(p, locs)
	loc := locs[[2]string{"read", "x"}]
	if loc.File != "p.rego" || loc.Line != 10 {
		t.Errorf("loc = %+v, want p.rego:10 (first occurrence)", loc)
	}
}
