//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-32 test — @ownership column must exist in DDL table

package ddl_rego

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestXdp32OwnershipColumn(t *testing.T) {
	if d := xdp32OwnershipColumn(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	fs := fsWithGroundTables([]string{"gigs"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "gig", Table: "gigs", Column: "owner_id", SourceLine: 3}, // present -> ok
			{Resource: "gig", Table: "gigs", Column: "missing", SourceLine: 4},  // absent -> diag
			{Resource: "x", Table: "ghost", Column: "c"},                        // table not in set -> skip (XDP-31)
			{Resource: "y", Table: "gigs", Column: ""},                          // empty col -> skip
		}},
	})
	fs.DDLTables = []ddl.Table{ddlTable("gigs", "owner_id")}

	d := xdp32OwnershipColumn(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDP-32]") || !strings.Contains(d[0].Message, "gigs.missing") {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
