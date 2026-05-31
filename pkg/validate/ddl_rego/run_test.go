//ff:func feature=validate type=test control=sequence topic=ddl-rego
//ff:what Run/helper test — XDP 규칙 일괄 실행 + buildDDLTableSet/ColumnIndex 검증
package ddl_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	// Empty Fullstack -> no diagnostics (each rule short-circuits).
	if d := Run(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty fs Run should yield no diags, got %v", d)
	}

	// One XDP-31 violation aggregated through Run.
	fs := fsWithGroundTables([]string{"other"}, []rego.Policy{
		{File: "p.rego", Ownerships: []rego.OwnershipMapping{
			{Resource: "gig", Table: "gigs", SourceLine: 1},
		}},
	})
	d := Run(fs)
	if len(d) != 1 {
		t.Fatalf("expected 1 aggregated diag, got %d: %v", len(d), d)
	}
}
