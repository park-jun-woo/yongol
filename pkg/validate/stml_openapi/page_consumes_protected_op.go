//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what pageConsumesProtectedOp — 페이지 소비 op(컴포넌트 포함) 중 security 보호 op 존재 여부 판정

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// pageConsumesProtectedOp reports whether the page consumes at least one
// security-protected OpenAPI operation — the per-page judgment
// react.resolveProtectedPages applies when flagging <ProtectedRoute> pages
// (PageConsumedOps consumption × OpRequiresAuth security inheritance).
// ops is the operationId universe (opMap key set) the component scan
// filters against.
func pageConsumesProtectedOp(p stml.PageSpec, fs *yongol.Fullstack, opMap map[string]operationEntry, ops map[string]struct{}) bool {
	for id := range PageConsumedOps(p, fs.SpecsDir, ops) {
		entry, ok := opMap[id]
		if !ok {
			continue
		}
		if OpRequiresAuth(entry.op, fs.OpenAPIDoc) {
			return true
		}
	}
	return false
}
