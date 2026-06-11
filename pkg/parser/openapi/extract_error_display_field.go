//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what ExtractErrorDisplayField — ErrorResponse 스키마에서 mutation onError가 읽을 표시 필드를 도출한다

package openapi

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// ExtractErrorDisplayField derives the field a generated mutation onError
// handler reads from a thrown server ErrorResponse. It scans schemas whose
// name contains "Error" (the XOE-01 heuristic) and returns:
//   - "error"   when a string property "error" exists (current ErrorResponse
//     contract — matches XOE-01's targets),
//   - "message" when "error" is absent but a string "message" exists,
//   - "error"   as the default when neither exists or no schema is present.
//
// The default keeps the generated handler's primary read schema-aligned and
// guarantees a non-empty field name, so codegen never emits a broken
// `e?. ?? e?.message` expression even for a nil document.
func ExtractErrorDisplayField(doc *openapi3.T) string {
	const defaultField = "error"
	if doc == nil || doc.Components == nil || doc.Components.Schemas == nil {
		return defaultField
	}

	names := make([]string, 0, len(doc.Components.Schemas))
	for name := range doc.Components.Schemas {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		if !strings.Contains(name, "Error") {
			continue
		}
		ref := doc.Components.Schemas[name]
		if ref == nil || ref.Value == nil {
			continue
		}
		if hasStringProperty(ref.Value, "error") {
			return "error"
		}
		if hasStringProperty(ref.Value, "message") {
			return "message"
		}
	}
	return defaultField
}
