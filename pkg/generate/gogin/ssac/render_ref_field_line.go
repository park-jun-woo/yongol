//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderRefFieldLine — $ref 타입 필드 ( convert<Type>* ) 응답 라인 렌더링

package ssac

import "fmt"

// renderRefFieldLine emits one field line for a $ref-typed response field.
// Four subcases (array×required) select the convert call shape:
//   required array    → convert<Type>List(v)
//   optional array    → ptrOf(convert<Type>List(v))
//   required scalar   → *convert<Type>(v)  (deref pointer)
//   optional scalar   → convert<Type>(v)   (leaves *Type)
func renderRefFieldLine(goFieldName, varExpr string, rf responseField) string {
	if rf.IsArray {
		if rf.IsRequired {
			return fmt.Sprintf("\t%s: convert%sList(%s),", goFieldName, rf.RefType, varExpr)
		}
		return fmt.Sprintf("\t%s: ptrOf(convert%sList(%s)),", goFieldName, rf.RefType, varExpr)
	}
	// convert<Type> always returns *api.Type. Deref for required.
	if rf.IsRequired {
		return fmt.Sprintf("\t%s: *convert%s(%s),", goFieldName, rf.RefType, varExpr)
	}
	return fmt.Sprintf("\t%s: convert%s(%s),", goFieldName, rf.RefType, varExpr)
}
