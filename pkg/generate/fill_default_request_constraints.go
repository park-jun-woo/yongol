//ff:func feature=generate type=util control=iteration dimension=1
//ff:what STML 폼 액션 중 RequestConstraints에 빠진 operationId를 OpenAPI 기본 타입으로 채운다
package generate

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// fillDefaultRequestConstraints ensures that every STML form action with fields
// has a corresponding entry in the RequestConstraints map. For operations that
// are missing (e.g. non-JSON content type or schemas that were skipped), it
// looks up the OpenAPI document and creates default FieldConstraint entries
// using only the property type and required flag.
//
// When the OpenAPI document does not contain a matching requestBody schema, a
// fallback entry is created from the STML field names with type "string".
//
// The function returns a new map; the original is not modified.
func fillDefaultRequestConstraints(
	pages []stmlparser.PageSpec,
	doc *openapi3.T,
	existing map[string]map[string]oapiparser.FieldConstraint,
) map[string]map[string]oapiparser.FieldConstraint {
	needed := collectFormActionOps(pages)
	if len(needed) == 0 {
		return existing
	}

	// Check which needed ops are already covered.
	var missing []actionEntry
	for _, ae := range needed {
		if fields, ok := existing[ae.opID]; ok && len(fields) > 0 {
			continue
		}
		missing = append(missing, ae)
	}
	if len(missing) == 0 {
		return existing
	}

	// Build augmented map (shallow copy + new entries).
	result := make(map[string]map[string]oapiparser.FieldConstraint, len(existing)+len(missing))
	for k, v := range existing {
		result[k] = v
	}
	for _, ae := range missing {
		if _, done := result[ae.opID]; done {
			continue
		}
		fields := resolveDefaultFields(ae, doc)
		if len(fields) > 0 {
			result[ae.opID] = fields
		}
	}
	return result
}
