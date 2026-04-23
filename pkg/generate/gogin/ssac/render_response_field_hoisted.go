//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.renderResponseFieldHoisted — hoist 된 convert 로컬 변수를 struct literal 에 주입

package ssac

import "fmt"

// renderResponseFieldHoisted is the variant of renderResponseField used
// by buildFieldResponse after convert<Model> became error-returning.
// $ref fields read from the pre-hoisted local variable (scalarLocal /
// listLocal); non-$ref fields delegate to the original rendering since
// their shape is unchanged.
func (g *methodGen) renderResponseFieldHoisted(
	jsonName, varExpr string,
	scalarLocal, listLocal map[string]string,
) string {
	goFieldName := pascalCase(jsonName)
	rf, hasSchema := g.RespFields[jsonName]

	if hasSchema && rf.RefType != "" {
		return renderRefResponseField(goFieldName, jsonName, rf, scalarLocal, listLocal)
	}
	if hasSchema && rf.IsRequired {
		return fmt.Sprintf("\t%s: %s,", goFieldName, varExpr)
	}
	if isLiteral(varExpr) {
		return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, varExpr)
	}
	return fmt.Sprintf("\t%s: &%s,", goFieldName, varExpr)
}
