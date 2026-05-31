//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-31 test — @ownership table must exist in DDL (Ground 기반)
package ddl_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdp31OwnershipTable(t *testing.T) {
	// nil Fullstack -> nil.
	if d := xdp31OwnershipTable(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	// No Ground -> nil (XDP-31 relies solely on Ground).
	noGround := &yongol.Fullstack{ParsedPolicies: []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{{Resource: "gig", Table: "gigs"}}},
	}}
	if d := xdp31OwnershipTable(noGround); d != nil {
		t.Errorf("no Ground should yield nil, got %v", d)
	}

	// Table present -> no diagnostics. Empty-table and duplicate are skipped.
	ok := fsWithGroundTables([]string{"gigs"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "gig", Table: "gigs", SourceLine: 3},
			{Resource: "gig", Table: "gigs", SourceLine: 3}, // duplicate -> seen
			{Resource: "x", Table: ""},                      // empty -> skip
		}},
	})
	if d := xdp31OwnershipTable(ok); len(d) != 0 {
		t.Errorf("expected no diags, got %v", d)
	}

	// Table missing -> one diagnostic.
	bad := fsWithGroundTables([]string{"other"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "gig", Table: "gigs", SourceLine: 7},
		}},
	})
	d := xdp31OwnershipTable(bad)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDP-31]") || d[0].Line != 7 {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
