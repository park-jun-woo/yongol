//ff:func feature=generate type=util control=sequence
//ff:what 누락된 operation에 대해 OpenAPI 조회 후 기본 FieldConstraint를 생성한다 (없으면 string 폴백)
package generate

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// resolveDefaultFields builds default FieldConstraint entries for a missing
// operation. It tries the OpenAPI doc first; falls back to string-typed entries
// derived from the STML field names.
func resolveDefaultFields(ae actionEntry, doc *openapi3.T) map[string]oapiparser.FieldConstraint {
	if fields := resolveFieldsFromOpenAPI(doc, ae.opID); len(fields) > 0 {
		return fields
	}
	return buildFallbackStringFields(ae.fieldNames)
}
