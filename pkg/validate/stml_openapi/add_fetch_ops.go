//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what addFetchOps — FetchBlock과 중첩 fetch의 operationId를 집합에 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// addFetchOps records f's operationId and recurses into its nested fetches.
func addFetchOps(f stml.FetchBlock, ops map[string]bool) {
	ops[f.OperationID] = true
	for _, child := range f.NestedFetches {
		addFetchOps(child, ops)
	}
}
