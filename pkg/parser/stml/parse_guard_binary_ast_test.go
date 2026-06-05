//ff:func feature=stml-parse type=test control=sequence dimension=1
//ff:what ParseGuard 이항(&&) AST 구조(좌/우 비교 노드) 검증

package stml

import "testing"

func TestParseGuardBinaryAST(t *testing.T) {
	expr, err := ParseGuard("workflow.status=active && currentUser.Role=owner")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expr.Kind != GuardBinary || expr.Op != "&&" {
		t.Fatalf("expected top-level && binary, got kind=%d op=%q", expr.Kind, expr.Op)
	}
	if expr.Left.Kind != GuardCompare || expr.Left.Ref.Path() != "workflow.status" || expr.Left.Value != "active" {
		t.Fatalf("unexpected left node: %+v", expr.Left)
	}
	if expr.Right.Kind != GuardCompare || expr.Right.Ref.Path() != "currentUser.Role" {
		t.Fatalf("unexpected right node: %+v", expr.Right)
	}
}
