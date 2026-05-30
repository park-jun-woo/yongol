//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-33 test — @ownership via join table must exist in DDL

package ddl_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestXdp33OwnershipJoinTable(t *testing.T) {
	if d := xdp33OwnershipJoinTable(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	fs := fsWithGroundTables([]string{"gigs"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "proposal", JoinTable: "gigs", SourceLine: 5}, // present -> ok
			{Resource: "proposal", JoinTable: "gigs", SourceLine: 5}, // dup -> seen
			{Resource: "bid", JoinTable: "missing_t", SourceLine: 8}, // absent -> diag
			{Resource: "x", JoinTable: ""},                           // empty -> skip
		}},
	})

	d := xdp33OwnershipJoinTable(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDP-33]") || !strings.Contains(d[0].Message, "missing_t") {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
