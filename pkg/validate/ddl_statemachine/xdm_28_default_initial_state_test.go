//ff:func feature=validate type=test control=sequence topic=states
//ff:what XDM-28/Run/helper test — 초기 전이 vs DDL DEFAULT 일치 + 대상 매핑 검증

package ddl_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDiagramDDLTarget(t *testing.T) {
	cases := []struct {
		id         string
		wantTable  string
		wantColumn string
	}{
		{"Course", "courses", "status"},      // PascalEntity convention
		{"order", "orders", "status"},        // no underscore
		{"order_phase", "orders", "phase"},   // entity_column convention
		{"_leading", "_leadings", "status"},  // idx<=0 -> status fallback
		{"trailing_", "trailing_", "status"}, // trailing underscore -> fallback (plural of whole lowercased id)
	}
	for _, c := range cases {
		gotT, gotC := diagramDDLTarget(c.id)
		if gotT != c.wantTable || gotC != c.wantColumn {
			t.Errorf("diagramDDLTarget(%q) = (%q,%q), want (%q,%q)", c.id, gotT, gotC, c.wantTable, c.wantColumn)
		}
	}
}

func TestCheckDefaultAgainstInitial(t *testing.T) {
	// Matching default -> nil.
	gMatch := &rule.Ground{Types: map[string]string{"DDL.default.value.orders.status": "draft"}}
	if d := checkDefaultAgainstInitial(gMatch, "order_status", "orders", "status", "draft"); d != nil {
		t.Errorf("matching default should yield nil, got %v", d)
	}

	// No default registered -> ERROR.
	gNone := &rule.Ground{Types: map[string]string{}}
	d1 := checkDefaultAgainstInitial(gNone, "order_status", "orders", "status", "draft")
	if d1 == nil || !strings.Contains(d1.Message, "has no DEFAULT") {
		t.Errorf("expected no-DEFAULT diag, got %v", d1)
	}

	// Mismatch -> ERROR.
	gMismatch := &rule.Ground{Types: map[string]string{"DDL.default.value.orders.status": "open"}}
	d2 := checkDefaultAgainstInitial(gMismatch, "order_status", "orders", "status", "draft")
	if d2 == nil || !strings.Contains(d2.Message, "≠") {
		t.Errorf("expected mismatch diag, got %v", d2)
	}
}

func TestXdm28DefaultInitialState(t *testing.T) {
	// No Ground -> nil.
	noG := &yongol.Fullstack{StateDiagrams: []*statemachine.StateDiagram{{ID: "order", InitialState: "draft"}}}
	if d := xdm28DefaultInitialState(noG); d != nil {
		t.Errorf("no Ground should yield nil, got %v", d)
	}

	g := &rule.Ground{Types: map[string]string{"DDL.default.value.orders.status": "draft"}}

	// InitialState empty skipped; nil skipped; matching diagram -> no diag.
	ok := fsWithGround(g,
		&statemachine.StateDiagram{ID: "order", InitialState: "draft"},
		&statemachine.StateDiagram{ID: "x", InitialState: ""},
		nil,
	)
	if d := xdm28DefaultInitialState(ok); len(d) != 0 {
		t.Errorf("expected no diags, got %v", d)
	}

	// Mismatch -> one diag.
	bad := fsWithGround(g, &statemachine.StateDiagram{ID: "order", InitialState: "open"})
	d := xdm28DefaultInitialState(bad)
	if len(d) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(d), d)
	}
	if !strings.Contains(d[0].Message, "[XDM-28]") {
		t.Errorf("unexpected diag: %+v", d[0])
	}
}

func TestRun(t *testing.T) {
	// No Ground -> both rules short-circuit -> nil.
	if d := Run(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("expected no diags for empty fs, got %v", d)
	}

	g := &rule.Ground{
		Lookup: map[string]rule.StringSet{"DDL.column.orders": {}},
		Types:  map[string]string{},
	}
	fs := fsWithGround(g, &statemachine.StateDiagram{ID: "order_status", InitialState: "draft"})
	d := Run(fs)
	// order_status missing column -> XDM-27 diag; missing default -> XDM-28 diag.
	if len(d) != 2 {
		t.Fatalf("expected 2 aggregated diags, got %d: %v", len(d), d)
	}
}
