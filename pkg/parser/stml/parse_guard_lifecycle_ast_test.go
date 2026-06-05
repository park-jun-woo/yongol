//ff:func feature=stml-parse type=test control=sequence dimension=1
//ff:what ParseGuard 그룹+생명주기 AST 구조 검증

package stml

import "testing"

func TestParseGuardLifecycleAST(t *testing.T) {
	expr, err := ParseGuard("(items.list.empty)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expr.Kind != GuardGroup {
		t.Fatalf("expected group, got kind=%d", expr.Kind)
	}
	if expr.Operand.Kind != GuardLifecycle || expr.Operand.Lifecycle != "empty" {
		t.Fatalf("unexpected lifecycle node: %+v", expr.Operand)
	}
}
