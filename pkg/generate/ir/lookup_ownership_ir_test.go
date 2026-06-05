//ff:func feature=gen-ir type=test control=sequence
//ff:what TestLookupOwnershipIR -- ParsedPolicies 순회로 resource 매칭 OwnershipInfo 조회 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestLookupOwnershipIR(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{Ownerships: []rego.OwnershipMapping{
				{Resource: "gig", Table: "gigs", Column: "client_id"},
			}},
			{Ownerships: []rego.OwnershipMapping{
				{Resource: "note", Table: "notes", Column: "owner_id"},
			}},
		},
	}

	// match found in the second policy
	info := lookupOwnershipIR(fs, "note")
	if info == nil {
		t.Fatal("expected OwnershipInfo for resource note, got nil")
	}
	if info.Table != "notes" || info.OwnerColumn != "owner_id" {
		t.Errorf("info = %+v, want Table=notes OwnerColumn=owner_id", info)
	}

	// no matching resource across all policies -> nil
	if got := lookupOwnershipIR(fs, "missing"); got != nil {
		t.Errorf("expected nil for unmatched resource, got %+v", got)
	}

	// no policies at all -> nil
	if got := lookupOwnershipIR(&yongol.Fullstack{}, "note"); got != nil {
		t.Errorf("expected nil with no policies, got %+v", got)
	}
}
