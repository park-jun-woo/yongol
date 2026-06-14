//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what resolveFallbackEntity — 비 SSaC 응답을 $ref 컴포넌트 또는 inline 컬럼매칭으로 엔티티 귀속

package openapi_ddl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// resolveFallbackEntity attributes a non-SSaC response to a canonical entity:
// a $ref response is accepted when its component is DDL-backed (entityComponent
// guard); an inline response defers to the column-set match (strategy B-2
// fallback). Returns "" when no entity can be attributed.
func resolveFallbackEntity(idx *entityIndex, schemaRef *openapi3.SchemaRef) string {
	if schemaRef.Ref != "" {
		name := schemaRef.Ref
		if i := strings.LastIndexByte(name, '/'); i >= 0 {
			name = name[i+1:]
		}
		return entityComponent(idx, name)
	}
	return inferInlineResponseEntity(idx, topLevelKeys(schemaRef))
}
