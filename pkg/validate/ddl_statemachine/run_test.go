//ff:func feature=validate type=test control=sequence topic=states
//ff:what XDM-28/Run/helper test — 초기 전이 vs DDL DEFAULT 일치 + 대상 매핑 검증
package ddl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
