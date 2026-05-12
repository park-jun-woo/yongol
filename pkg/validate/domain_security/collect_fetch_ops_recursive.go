//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what collectFetchOpsRecursive — FetchBlock에서 재귀적으로 operationId 수집
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// collectFetchOpsRecursive recursively collects operationIds from a FetchBlock.
func collectFetchOpsRecursive(f stml.FetchBlock, out map[string]struct{}) {
	out[f.OperationID] = struct{}{}
	for _, child := range f.NestedFetches {
		collectFetchOpsRecursive(child, out)
	}
}
