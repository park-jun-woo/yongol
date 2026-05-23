//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what FetchBlock 의 Eaches·NestedFetches 에서 진단을 재귀 수집
package stml

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func collectFetchDiags(fb *FetchBlock, file string, out *[]diagnostic.Diagnostic) {
	for i := range fb.Eaches {
		appendEachDiags(&fb.Eaches[i], file, out)
	}
	for i := range fb.NestedFetches {
		collectFetchDiags(&fb.NestedFetches[i], file, out)
	}
}
