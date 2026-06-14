//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what resolveSSaCEntity — SSaC @response 로부터 canonical component 해석 (전략 A: Target / B-1: Fields)

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// resolveSSaCEntity resolves the canonical entity component from an SSaC
// @response sequence. Strategy A handles shorthand `@response <var>` (rejecting
// collection-bound vars); strategy B-1 handles `@response {k: v}`. Both feed the
// B-2 DDL guard (entityComponent). Returns "" for non-entity responses.
func resolveSSaCEntity(idx *entityIndex, fn *ssac.ServiceFunc, seq *ssac.Sequence) string {
	if seq.Target != "" {
		if varBindingIsCollection(fn, seq.Target) {
			return ""
		}
		raw := idx.g.Types["SSaC.var."+fn.Name+"."+seq.Target]
		if raw == "" || isCollectionType(raw) {
			return ""
		}
		return entityComponent(idx, normalizeTypeName(raw))
	}
	if len(seq.Fields) > 0 {
		return entityComponent(idx, inferFieldsResponseEntity(idx, fn, seq.Fields))
	}
	return ""
}
