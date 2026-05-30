//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func authSeq(action, resource string, line int) ssac.Sequence {
	return ssac.Sequence{Type: "auth", Action: action, Resource: resource, Line: line}
}

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

func TestCollectSSaCAuthPairs(t *testing.T) {
	funcs := []ssac.ServiceFunc{{
		Sequences: []ssac.Sequence{
			authSeq("delete", "project", 5),
			{Type: "get"}, // non-auth ignored
		},
	}}
	pairs := collectSSaCAuthPairs(funcs)
	if !pairs[[2]string{"delete", "project"}] {
		t.Errorf("pairs = %v", pairs)
	}
	if len(pairs) != 1 {
		t.Errorf("expected only one pair, got %v", pairs)
	}
}

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

func TestScanSSaCAuthLocations(t *testing.T) {
	locs := map[[2]string]PairLocation{}
	fn := ssac.ServiceFunc{
		FileName: "svc.ssac",
		Sequences: []ssac.Sequence{
			authSeq("delete", "project", 3),
			authSeq("delete", "project", 9), // dup
		},
	}
	scanSSaCAuthLocations(fn, locs)
	loc := locs[[2]string{"delete", "project"}]
	if loc.File != "svc.ssac" || loc.Line != 3 {
		t.Errorf("loc = %+v, want svc.ssac:3", loc)
	}
}

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

func TestFirstSSaCAuthLocation(t *testing.T) {
	funcs := []ssac.ServiceFunc{{FileName: "s.ssac", Sequences: []ssac.Sequence{authSeq("write", "doc", 4)}}}
	locs := firstSSaCAuthLocation(funcs)
	if locs[[2]string{"write", "doc"}].Line != 4 {
		t.Errorf("locs = %+v", locs)
	}
}

func TestXps28MissingRegoDiag(t *testing.T) {
	pair := [2]string{"delete", "project"}
	pairLoc := map[[2]string]PairLocation{pair: {File: "s.ssac", Line: 12}}

	// Pair absent from rego → diagnostic.
	diag, ok := xps28MissingRegoDiag(pair, map[[2]string]bool{}, pairLoc)
	if !ok {
		t.Fatal("expected diagnostic for missing rego rule")
	}
	if diag.Line != 12 || diag.OperationID != "delete" {
		t.Errorf("diag = %+v", diag)
	}
	// Pair present → no diagnostic.
	if _, ok := xps28MissingRegoDiag(pair, map[[2]string]bool{pair: true}, pairLoc); ok {
		t.Error("present pair should not produce a diagnostic")
	}
}

func TestXsp29MissingSSaCDiag(t *testing.T) {
	pair := [2]string{"read", "doc"}
	pairLoc := map[[2]string]PairLocation{pair: {File: "p.rego", Line: 8}}

	diag, ok := xsp29MissingSSaCDiag(pair, map[[2]string]bool{}, pairLoc)
	if !ok {
		t.Fatal("expected diagnostic for missing SSaC @auth")
	}
	if diag.Line != 8 || diag.File != "p.rego" {
		t.Errorf("diag = %+v", diag)
	}
	if _, ok := xsp29MissingSSaCDiag(pair, map[[2]string]bool{pair: true}, pairLoc); ok {
		t.Error("present pair should not produce a diagnostic")
	}
}
