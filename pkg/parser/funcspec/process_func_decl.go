//ff:func feature=funcspec type=parser control=sequence
//ff:what processFuncDecl — populate FuncSpec from the matching FuncDecl (HasBody, ReturnTypes)

package funcspec

import (
	"go/ast"
	"go/token"
)

// processFuncDecl populates the parts of FuncSpec that come from the actual
// function declaration: HasBody (stub detection) and ReturnTypes (signature).
// Only fires for the FuncDecl whose name matches `ucFirst(spec.Name)`.
func processFuncDecl(d *ast.FuncDecl, fset *token.FileSet, spec *FuncSpec) {
	funcName := ucFirst(spec.Name)
	if d.Name.Name != funcName {
		return
	}
	if d.Body != nil {
		spec.HasBody = !isStubBody(fset, d.Body)
	}
	spec.ReturnTypes = extractReturnTypes(fset, d)
}
