//ff:func feature=generate type=util control=iteration dimension=1
//ff:what 필드 이름 목록에서 string 타입의 폴백 FieldConstraint를 생성한다
package generate

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// buildFallbackStringFields creates string-typed FieldConstraint entries from
// the given field names. Used when an operation is not found in the OpenAPI doc.
func buildFallbackStringFields(fieldNames []string) map[string]oapiparser.FieldConstraint {
	fields := make(map[string]oapiparser.FieldConstraint, len(fieldNames))
	for _, name := range fieldNames {
		fields[name] = oapiparser.FieldConstraint{Type: "string"}
	}
	return fields
}
