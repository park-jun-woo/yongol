//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.renderResponseField — 한 필드에 대한 typed response 라인 생성 ($ref / 필수 / 리터럴 분기)

package ssac

import "fmt"

// renderResponseField produces one field-assignment line for a typed 200
// response. Four mutually exclusive cases drive the output shape:
//   - $ref type (array or scalar, required vs optional) → convert<Type>() call
//   - required primitive                                → direct value
//   - literal expression (isLiteral)                    → ptrOf(literal)
//   - anything else                                     → &variable
func (g *methodGen) renderResponseField(jsonName, varExpr string) string {
	goFieldName := pascalCase(jsonName)
	rf, hasSchema := g.RespFields[jsonName]

	if hasSchema && rf.RefType != "" {
		return renderRefFieldLine(goFieldName, varExpr, rf)
	}
	if hasSchema && rf.IsRequired {
		// Required primitive → direct assignment, no pointer.
		return fmt.Sprintf("\t%s: %s,", goFieldName, varExpr)
	}
	if isLiteral(varExpr) {
		// literal (true, false, number, "string") → use ptrOf()
		return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, varExpr)
	}
	// variable → address-of
	return fmt.Sprintf("\t%s: &%s,", goFieldName, varExpr)
}
