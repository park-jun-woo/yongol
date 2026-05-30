//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-34 test — @ownership via join FK column must exist in DDL join table

package ddl_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestXdp34OwnershipJoinColumn(t *testing.T) {
	if d := xdp34OwnershipJoinColumn(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	fs := fsWithGroundTables([]string{"gigs"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "proposal", JoinTable: "gigs", JoinFK: "client_id", SourceLine: 5}, // present -> ok
			{Resource: "bid", JoinTable: "gigs", JoinFK: "missing_fk", SourceLine: 8},     // absent -> diag
			{Resource: "x", JoinTable: "ghost", JoinFK: "fk"},                             // table not in set -> skip
			{Resource: "y", JoinTable: "gigs", JoinFK: ""},                                // empty fk -> skip
		}},
	})
	fs.DDLTables = []ddl.Table{ddlTable("gigs", "client_id")}

	d := xdp34OwnershipJoinColumn(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDP-34]") || !strings.Contains(d[0].Message, "gigs.missing_fk") {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
