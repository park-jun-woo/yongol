//ff:func feature=funcspec type=test control=sequence
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine
package funcspec

import (
	"testing"
)

func TestFindFuncDeclLine(t *testing.T) {
	fset, f := parseDeclT(t, "package p\n\nfunc Target() {}\n")
	if line := findFuncDeclLine(f, fset, "Target"); line != 3 {
		t.Errorf("line = %d, want 3", line)
	}
	if line := findFuncDeclLine(f, fset, "Missing"); line != 0 {
		t.Errorf("missing func line = %d, want 0", line)
	}
}
