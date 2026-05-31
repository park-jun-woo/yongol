//ff:func feature=validate type=test control=sequence topic=states
//ff:what XDM-28/Run/helper test — 초기 전이 vs DDL DEFAULT 일치 + 대상 매핑 검증
package ddl_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

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
