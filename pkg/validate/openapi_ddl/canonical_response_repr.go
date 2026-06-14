//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-11/XDO-12 — 같은 엔티티의 2xx 응답이 동일 표현(canonical component)으로 수렴하는지 검증

package openapi_ddl

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// canonicalResponseRepr validates "one resource = one representation":
//
//	XDO-11 (ERROR)   — two or more 2xx responses returning the same entity
//	                   expose different representations (field set / nesting /
//	                   flat-vs-wrapper). Direct contract contradiction (BUG-131).
//	XDO-12 (WARNING) — an entity response is defined inline without sharing a
//	                   component $ref; consistent today but drift-prone.
//
// Entity identity is resolved via strategy A (SSaC @response var), B-1 (SSaC
// @response field base-var convergence) and B-2 (DDL table/component guard +
// non-SSaC column-set fallback). Non-entity responses are out of scope.
func canonicalResponseRepr(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.OpenAPIDoc.Components.Schemas == nil {
		return nil
	}
	if fs.Ground() == nil {
		return nil
	}
	idx := buildEntityIndex(fs)
	groups := collectEntityResponses(fs, idx)

	comps := make([]string, 0, len(groups))
	for comp := range groups {
		comps = append(comps, comp)
	}
	sort.Strings(comps)

	var diags []diagnostic.Diagnostic
	for _, comp := range comps {
		diags = append(diags, evalEntityGroup(comp, groups[comp])...)
	}
	return diags
}
