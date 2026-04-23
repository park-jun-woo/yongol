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
		if rf.IsArray {
			local := listLocal[jsonName]
			// convert<Type>List returns ([]api.Type, error) — local is
			// already the slice value. Required fields expect []api.Type;
			// optional want a pointer.
			if rf.IsRequired {
				return fmt.Sprintf("\t%s: %s,", goFieldName, local)
			}
			return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, local)
		}
		local := scalarLocal[jsonName]
		// convert<Type> returns (*api.Type, error) — local is *api.Type.
		// Required fields want api.Type (deref); optional fields want
		// *api.Type (as-is).
		if rf.IsRequired {
			return fmt.Sprintf("\t%s: *%s,", goFieldName, local)
		}
		return fmt.Sprintf("\t%s: %s,", goFieldName, local)
	}
	if hasSchema && rf.IsRequired {
		return fmt.Sprintf("\t%s: %s,", goFieldName, varExpr)
	}
	if isLiteral(varExpr) {
		return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, varExpr)
	}
	return fmt.Sprintf("\t%s: &%s,", goFieldName, varExpr)
}
