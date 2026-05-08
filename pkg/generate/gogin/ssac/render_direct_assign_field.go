//ff:func feature=gen-gogin type=util control=selection topic=response
//ff:what methodGen.renderDirectAssignField — @call 결과 변수를 converter 없이 직접 대입하는 응답 필드 렌더

package ssac

import "fmt"

// renderDirectAssignField renders a response struct field that references
// a @call result variable. Func Response types are user-authored and
// OpenAPI-compatible, so no convert<Type>() wrapper is needed (BUG-050).
// Required fields use value assignment; optional fields take the address.
func (g *methodGen) renderDirectAssignField(jsonName, varExpr string) string {
	goFieldName := pascalCase(jsonName)
	rf := g.RespFields[jsonName]
	if rf.IsRequired {
		return fmt.Sprintf("\t%s: %s,", goFieldName, varExpr)
	}
	return fmt.Sprintf("\t%s: &%s,", goFieldName, varExpr)
}
