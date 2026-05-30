//ff:func feature=validate type=test control=sequence topic=states
//ff:what XDM-27 test — stateDiagram <entity>_<field> → DDL column 존재 검증

package ddl_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// fsWithGround builds a Fullstack carrying the given Ground and state diagrams.
func fsWithGround(g *rule.Ground, diagrams ...*statemachine.StateDiagram) *yongol.Fullstack {
	fs := &yongol.Fullstack{StateDiagrams: diagrams}
	fs.SetGround(g)
	return fs
}

func TestXdm27StateFieldColumn(t *testing.T) {
	// No Ground -> nil.
	noG := &yongol.Fullstack{StateDiagrams: []*statemachine.StateDiagram{{ID: "order_status"}}}
	if d := xdm27StateFieldColumn(noG); d != nil {
		t.Errorf("no Ground should yield nil, got %v", d)
	}

	g := &rule.Ground{Lookup: map[string]rule.StringSet{
		"DDL.column.orders": {"status": true},
	}}

	// Column present (orders.status) -> no diag; nil diagram skipped;
	// underscore-less ID skipped.
	ok := fsWithGround(g,
		&statemachine.StateDiagram{ID: "order_status"},
		nil,
		&statemachine.StateDiagram{ID: "Course"},
	)
	if d := xdm27StateFieldColumn(ok); len(d) != 0 {
		t.Errorf("expected no diags, got %v", d)
	}

	// Missing column -> one diag.
	bad := fsWithGround(g, &statemachine.StateDiagram{ID: "order_phase"})
	d := xdm27StateFieldColumn(bad)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDM-27]") || !strings.Contains(d[0].Message, "orders.phase") {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}
