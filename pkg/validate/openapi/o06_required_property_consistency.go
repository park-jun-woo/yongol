//ff:func feature=validate type=rule control=iteration dimension=1 topic=openapi-structural
//ff:what O-6 — 모든 스키마(components+inline)의 required 항목이 properties에 선언돼야 함, 위반 시 ERROR

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o06RequiredPropertyConsistency validates O-6: every entry of a schema's
// `required` array must be declared in that schema's `properties`. Applies to
// components.schemas plus every operation requestBody/response inline schema and
// their nested schemas. A dangling required (name present in required but absent
// from properties) is an internal OpenAPI contradiction — oapi-codegen ignores
// it or emits a wrong stub — so it is reported as ERROR. This closes the false
// negative documented in BUG-105.
func o06RequiredPropertyConsistency(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, entry := range o06CollectAllSchemas(fs) {
		diags = append(diags, o06CheckSchemaRequired(fs, entry)...)
	}
	return diags
}
