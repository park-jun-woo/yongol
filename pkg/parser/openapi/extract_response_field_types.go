//ff:func feature=openapi-parse type=parser control=iteration dimension=2
//ff:what ExtractResponseFieldTypes — 응답 스키마의 필드 경로별 타입/포맷을 operationId 별로 추출한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractResponseFieldTypes returns a map of operationId → field path →
// FieldTypeInfo{Type, Format} from each operation's 2xx response schema. The
// keyed paths mirror the data-bind names the react emitter renders:
//
//   - top-level scalar/object properties (e.g. "can_delete", "summary")
//   - object dotted paths one level deep (e.g. "summary.credits_balance")
//   - array item properties (e.g. "photos.url")
//
// Unlike ResponseArrayItemTypes (which carries only item field types for the
// Number() wrap of row-action arguments), this map covers the full bind
// surface and includes Format so the emitter can branch on date/date-time
// (plans/gen/frontend Phase037, BUG-126). allOf members are resolved one
// level deep, matching the validate-side responseFields judgment.
func ExtractResponseFieldTypes(doc *openapi3.T) map[string]map[string]FieldTypeInfo {
	result := make(map[string]map[string]FieldTypeInfo)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			collectResponseFieldTypesForOp(result, op)
		}
	}
	return result
}
