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
