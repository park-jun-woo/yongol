//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestFlattenBlockStmts — BlockStmt 트리 하위 stmt 단일 슬라이스 수집 검증
package contract

import (
	"testing"
)

func TestFlattenBlockStmts(t *testing.T) {
	// Block with a nested if and a closure; the closure body's stmt
	// must NOT be flattened in.
	block := mustBlock(t, "x := 1\nif x > 0 { y := 2; _ = y }\ndefer func() { z := 3; _ = z }()\n_ = x\n")
	flat := flattenBlockStmts(block)
	if len(flat) == 0 {
		t.Fatal("expected non-empty flattened stmts")
	}
	// The closure's inner assignment (z := 3) must NOT leak in.
	if assignIdentPresent(flat, "z") {
		t.Fatalf("closure body stmt leaked into flatten output")
	}
	// The nested-if assignment (y := 2) SHOULD be present.
	if !assignIdentPresent(flat, "y") {
		t.Fatal("nested-if assignment should be flattened in")
	}
}
