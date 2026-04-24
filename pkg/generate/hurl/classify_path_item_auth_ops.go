//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what classifyPathItemAuthOps — 단일 PathItem 의 operation 들을 auth shape 로 분류

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// classifyPathItemAuthOps iterates the operations on a single OpenAPI
// PathItem and appends matching candidates to the signup / login slices.
// Ambiguous or non-matching shapes yield warnings via the provided slice
// pointer. Extracted from detectAuthOps to keep nesting depth within Q1
// dimension=2 (path iterate → classify; classify iterates methods).
func classifyPathItemAuthOps(
	path string,
	pathItem *openapi3.PathItem,
	funcsByOpID map[string]*ssac.ServiceFunc,
	signupCands *[]detectedAuthOp,
	loginCands *[]detectedAuthOp,
	warnings *[]string,
) {
	for method, op := range pathItem.Operations() {
		if op == nil || op.OperationID == "" {
			continue
		}
		if !isPublicOp(op) {
			continue
		}
		if !hasPasswordField(op) {
			continue
		}
		fn := funcsByOpID[op.OperationID]
		classifyAuthOpShape(path, method, op, fn, signupCands, loginCands, warnings)
	}
}
