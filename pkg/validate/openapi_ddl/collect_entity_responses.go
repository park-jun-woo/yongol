//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what collectEntityResponses — 모든 operation 의 2xx 응답을 canonical 엔티티별로 그룹핑

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectEntityResponses groups every operation's 2xx JSON response by its
// canonical entity component (resolveResponseEntity). Non-entity responses
// (scalars, lists, ad-hoc projections) are skipped. The returned map keys are
// component names; values are per-operation representation signatures.
func collectEntityResponses(fs *yongol.Fullstack, idx *entityIndex) map[string][]entityRepr {
	groups := make(map[string][]entityRepr)
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return groups
	}
	for _, item := range fs.OpenAPIDoc.Paths.Map() {
		for _, ve := range []struct {
			method string
			op     *openapi3.Operation
		}{
			{"GET", item.Get},
			{"POST", item.Post},
			{"PUT", item.Put},
			{"DELETE", item.Delete},
			{"PATCH", item.Patch},
		} {
			op := ve.op
			if op == nil || op.OperationID == "" {
				continue
			}
			schemaRef := successResponseSchemaRef(op, ve.method)
			if schemaRef == nil {
				continue
			}
			comp := resolveResponseEntity(idx, idx.funcByName[op.OperationID], schemaRef)
			if comp == "" {
				continue
			}
			groups[comp] = append(groups[comp], entityRepr{
				opID:     op.OperationID,
				line:     fs.OpenAPILines.OperationLine(op.OperationID),
				shapeKey: responseShapeKey(schemaRef),
			})
		}
	}
	return groups
}
