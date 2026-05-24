//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderPgtypeField — pgtype 변환된 필드의 struct literal 라인 생성 (required/optional 분기)

package ssac

import "fmt"

// renderPgtypeField emits the struct literal line for a pgtype-converted field.
// The pgtypex bridge call already returns the correct Go type — nullable columns
// yield pointer types (no extra wrapping), NOT NULL pgtype columns yield value
// types (wrap with ptrOf when the API field is optional).
func (g *methodGen) renderPgtypeField(jsonName, varExpr, conv string) string {
	goFieldName := pascalCase(jsonName)
	rf, hasSchema := g.RespFields[jsonName]
	if hasSchema && rf.IsRequired {
		return fmt.Sprintf("\t%s: %s,", goFieldName, conv)
	}
	// The pgtypex Ptr variant already returns *T for nullable
	// columns — assign directly. For NOT NULL pgtype (value
	// type), the ConvertExpr uses the non-Ptr variant which
	// returns a value; check via isAlreadyPointerApiField.
	if g.isPgtypeAlreadyPointer(varExpr) {
		return fmt.Sprintf("\t%s: %s,", goFieldName, conv)
	}
	return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, conv)
}
