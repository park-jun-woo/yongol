//ff:func feature=stml-parse type=parser control=sequence
//ff:what ChildNode 의 Fetch·Each 에서 진단을 수집
package stml

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func collectChildDiags(cn *ChildNode, file string, out *[]diagnostic.Diagnostic) {
	if cn.Fetch != nil {
		collectFetchDiags(cn.Fetch, file, out)
	}
	if cn.Each != nil {
		appendEachDiags(cn.Each, file, out)
	}
}
