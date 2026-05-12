//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what collectConsumedOpsFromPages — STML 페이지에서 소비된 operationId 수집
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// collectConsumedOpsFromPages collects operationIds from fetch and action blocks.
func collectConsumedOpsFromPages(pages []stml.PageSpec) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		for _, f := range page.Fetches {
			collectFetchOpsRecursive(f, out)
		}
		for _, a := range page.Actions {
			out[a.OperationID] = struct{}{}
		}
	}
	return out
}
