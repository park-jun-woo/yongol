//ff:func feature=orchestrator type=loader control=iteration dimension=1
//ff:what 한 도메인 doc 의 request/response 제약을 추출해 전역 fs.RequestConstraints/ResponseConstraints 에 병합 (opID 전역 유일·XDO-90)
package yongol

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// mergeDomainConstraints extracts request/response field constraints from one
// domain's OpenAPI doc and merges them into the global fs.RequestConstraints /
// fs.ResponseConstraints maps, lazily initializing them when nil. operationIds
// are globally unique across domains (XDO-90), so a plain per-opID map merge is
// sufficient — no key collisions. Domain-mode only: the single-site path is
// populated by parseOpenAPIIfPresent and is never routed through here.
func mergeDomainConstraints(fs *Fullstack, doc *openapi3.T, lines *oapiparser.LineIndex) {
	if fs.RequestConstraints == nil {
		fs.RequestConstraints = make(map[string]map[string]oapiparser.FieldConstraint)
	}
	for opID, fields := range oapiparser.ExtractRequestConstraints(doc, lines) {
		fs.RequestConstraints[opID] = fields
	}
	if fs.ResponseConstraints == nil {
		fs.ResponseConstraints = make(map[string]map[string]oapiparser.FieldConstraint)
	}
	for opID, fields := range oapiparser.ExtractResponseConstraints(doc, lines) {
		fs.ResponseConstraints[opID] = fields
	}
}
