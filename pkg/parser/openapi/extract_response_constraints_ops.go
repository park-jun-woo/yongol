//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractResponseConstraintsOps — 단일 PathItem의 operation별 response 제약조건 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

func extractResponseConstraintsOps(result map[string]map[string]FieldConstraint, item *openapi3.PathItem, lines *LineIndex) {
	for _, op := range item.Operations() {
		collectResponseConstraintsForOp(result, op, lines)
	}
}
