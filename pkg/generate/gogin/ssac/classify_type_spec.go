//ff:func feature=gen-gogin type=util control=selection
//ff:what classifyTypeSpec — 단일 TypeSpec를 embedded struct / schema alias로 판별해 respShape 반환

package ssac

import "go/ast"

// classifyTypeSpec maps a single TypeSpec to a respShape. It recognises the
// embedded form `struct{ <Ident> }` (single anonymous embedded field) and the
// alias form `<Ident>`. Returns ok=false for shapes that are neither (which
// fall back to alias at emit time).
func classifyTypeSpec(ts *ast.TypeSpec) (respShape, bool) {
	switch t := ts.Type.(type) {
	case *ast.StructType:
		if t.Fields == nil || len(t.Fields.List) != 1 {
			return respShape{}, false
		}
		field := t.Fields.List[0]
		if len(field.Names) != 0 {
			return respShape{}, false // not an anonymous embed
		}
		ident, ok := field.Type.(*ast.Ident)
		if !ok {
			return respShape{}, false
		}
		return respShape{Kind: shapeEmbedded, EmbeddedType: ident.Name}, true
	case *ast.Ident:
		return respShape{Kind: shapeAlias}, true
	default:
		return respShape{}, false
	}
}
