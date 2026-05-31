//ff:func feature=funcspec type=test control=sequence
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine
package funcspec

import (
	"testing"
)

func TestProcessFuncDecl(t *testing.T) {
	t.Run("matching name with body", func(t *testing.T) {
		fset, f := parseDeclT(t, `package p
func HashPassword() (Resp, error) { return Resp{Status: "x"}, nil }`)
		spec := &FuncSpec{Name: "hash_password"}
		processFuncDecl(firstFunc(f), fset, spec)
		if !spec.HasBody {
			t.Errorf("expected HasBody true")
		}
		if len(spec.ReturnTypes) != 2 {
			t.Errorf("ReturnTypes = %v", spec.ReturnTypes)
		}
	})
	t.Run("name mismatch is no-op", func(t *testing.T) {
		fset, f := parseDeclT(t, "package p\nfunc Other() {}")
		spec := &FuncSpec{Name: "hash_password"}
		processFuncDecl(firstFunc(f), fset, spec)
		if spec.HasBody || spec.ReturnTypes != nil {
			t.Errorf("expected no-op, got %+v", spec)
		}
	})
}
