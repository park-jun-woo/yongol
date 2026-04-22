//ff:func feature=gen-splitter type=util control=sequence
//ff:what genDeclDoc — GenDecl 또는 첫 TypeSpec 의 doc comment 텍스트
package splitter

import "go/ast"

// genDeclDoc returns the doc comment for a GenDecl. oapi-codegen often
// attaches the documentation to the inner TypeSpec rather than the
// GenDecl wrapper; we check the wrapper first, then the inner spec.
func genDeclDoc(d *ast.GenDecl) string {
	if d.Doc != nil {
		return d.Doc.Text()
	}
	if len(d.Specs) != 1 {
		return ""
	}
	ts, ok := d.Specs[0].(*ast.TypeSpec)
	if !ok || ts.Doc == nil {
		return ""
	}
	return ts.Doc.Text()
}
