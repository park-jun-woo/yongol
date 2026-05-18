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

func collectFetchDiags(fb *FetchBlock, file string, out *[]diagnostic.Diagnostic) {
	for i := range fb.Eaches {
		appendEachDiags(&fb.Eaches[i], file, out)
	}
	for i := range fb.NestedFetches {
		collectFetchDiags(&fb.NestedFetches[i], file, out)
	}
}

func collectChildDiags(cn *ChildNode, file string, out *[]diagnostic.Diagnostic) {
	if cn.Fetch != nil {
		collectFetchDiags(cn.Fetch, file, out)
	}
	if cn.Each != nil {
		appendEachDiags(cn.Each, file, out)
	}
}

func appendEachDiags(eb *EachBlock, file string, out *[]diagnostic.Diagnostic) {
	for _, d := range eb.Diags {
		if d.File == "" {
			d.File = file
		}
		*out = append(*out, d)
	}
}
