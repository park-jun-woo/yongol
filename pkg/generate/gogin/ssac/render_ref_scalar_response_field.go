//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderRefScalarResponseField — $ref 스칼라 응답 필드의 struct literal 라인 (required/optional 분기)

package ssac

import "fmt"

// convert<Type> returns (*api.Type, error) — local is *api.Type.
// Required fields want api.Type (deref); optional fields want
// *api.Type (as-is).
func renderRefScalarResponseField(goFieldName, jsonName string, rf responseField, scalarLocal map[string]string) string {
	local := scalarLocal[jsonName]
	if rf.IsRequired {
		return fmt.Sprintf("\t%s: *%s,", goFieldName, local)
	}
	return fmt.Sprintf("\t%s: %s,", goFieldName, local)
}
