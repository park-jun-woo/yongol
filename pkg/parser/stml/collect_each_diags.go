//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what PageSpec 내 모든 EachBlock에서 진단을 수집
package stml

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// collectEachDiags gathers diagnostics from every EachBlock reachable
// from the given PageSpec (through Fetches, NestedFetches, and Children).
func collectEachDiags(page *PageSpec, file string) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for i := range page.Fetches {
		collectFetchDiags(&page.Fetches[i], file, &out)
	}
	for i := range page.Children {
		collectChildDiags(&page.Children[i], file, &out)
	}
	return out
}
