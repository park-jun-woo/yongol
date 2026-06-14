//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what responseShapeKey — 응답 표현의 동치 비교용 시그니처 생성 ($ref 또는 inline 필드집합)

package openapi_ddl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// responseShapeKey produces the canonical-equivalence signature of a 2xx
// response. Per Phase005's resolution (i): two responses are the same
// representation iff they $ref the same component, or are inline with the same
// top-level field set. A "ref:<Component>" key never equals an "inline:<fields>"
// key, so a flat-inline vs $ref divergence (BUG-131) is detected.
func responseShapeKey(schemaRef *openapi3.SchemaRef) string {
	if schemaRef.Ref != "" {
		name := schemaRef.Ref
		if i := strings.LastIndexByte(name, '/'); i >= 0 {
			name = name[i+1:]
		}
		return "ref:" + name
	}
	return "inline:" + strings.Join(topLevelKeys(schemaRef), ",")
}
