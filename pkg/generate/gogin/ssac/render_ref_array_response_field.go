//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderRefArrayResponseField — $ref 배열 응답 필드의 struct literal 라인 (required/optional 분기)

package ssac

import "fmt"

// convert<Type>List returns ([]api.Type, error) — local is
// already the slice value. Required fields expect []api.Type;
// optional want a pointer.
func renderRefArrayResponseField(goFieldName, jsonName string, rf responseField, listLocal map[string]string) string {
	local := listLocal[jsonName]
	if rf.IsRequired {
		return fmt.Sprintf("\t%s: %s,", goFieldName, local)
	}
	return fmt.Sprintf("\t%s: ptrOf(%s),", goFieldName, local)
}
