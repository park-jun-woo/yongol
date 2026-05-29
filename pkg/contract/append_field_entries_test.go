//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestAppendFieldEntries — 다중 이름 그룹 분해 및 익명 필드 단일 엔트리 분기 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestAppendFieldEntries(t *testing.T) {
	// Multi-name group `a, b int` expands to two entries sharing type.
	named := &ast.Field{
		Names: []*ast.Ident{{Name: "a"}, {Name: "b"}},
	}
	out := appendFieldEntries(nil, named, "int", true)
	if len(out) != 2 {
		t.Fatalf("named: got %d entries want 2 (%v)", len(out), out)
	}
	if out[0] != (FuncParam{Name: "a", Type: "int"}) || out[1] != (FuncParam{Name: "b", Type: "int"}) {
		t.Errorf("named: got %+v", out)
	}

	// Anonymous field (no names) yields a single Name=="" entry.
	anon := &ast.Field{}
	out2 := appendFieldEntries(nil, anon, "error", false)
	if len(out2) != 1 {
		t.Fatalf("anon: got %d entries want 1 (%v)", len(out2), out2)
	}
	if out2[0] != (FuncParam{Name: "", Type: "error"}) {
		t.Errorf("anon: got %+v want {Name: empty, Type: error}", out2[0])
	}

	// Appends onto an existing slice rather than replacing.
	pre := []FuncParam{{Name: "x", Type: "bool"}}
	out3 := appendFieldEntries(pre, anon, "string", true)
	if len(out3) != 2 || out3[0].Name != "x" {
		t.Errorf("append: got %+v", out3)
	}
}
