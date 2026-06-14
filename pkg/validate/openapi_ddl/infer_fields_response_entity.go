//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what inferFieldsResponseEntity — @response {k:v} 값들의 base var 를 단일 모델로 수렴 (전략 B-1)

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// inferFieldsResponseEntity implements strategy B-1: collect each @response
// field value's base var, dereference it through Ground's SSaC.var type map
// (the same lookup XOS-67 uses), and converge to a single dominant model name.
// Literals are ignored; a collection-bound var, an unresolvable var, or a mix
// of distinct models yields "" (non-entity / envelope). The returned model is
// only a candidate — the caller still applies the B-2 DDL guard.
func inferFieldsResponseEntity(idx *entityIndex, fn *ssac.ServiceFunc, fields map[string]string) string {
	converged := ""
	for _, value := range fields {
		base := responseValueBaseVar(value)
		if base == "" {
			continue
		}
		if varBindingIsCollection(fn, base) {
			return ""
		}
		raw := idx.g.Types["SSaC.var."+fn.Name+"."+base]
		if raw == "" || isCollectionType(raw) {
			return ""
		}
		model := normalizeTypeName(raw)
		if converged == "" {
			converged = model
		} else if converged != model {
			return ""
		}
	}
	return converged
}
