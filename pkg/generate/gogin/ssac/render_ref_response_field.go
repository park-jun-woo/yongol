//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderRefResponseField — $ref 타입 응답 필드의 struct literal 라인 렌더 (array/scalar 분기)

package ssac

func renderRefResponseField(goFieldName, jsonName string, rf responseField, scalarLocal, listLocal map[string]string) string {
	if rf.IsArray {
		return renderRefArrayResponseField(goFieldName, jsonName, rf, listLocal)
	}
	return renderRefScalarResponseField(goFieldName, jsonName, rf, scalarLocal)
}
