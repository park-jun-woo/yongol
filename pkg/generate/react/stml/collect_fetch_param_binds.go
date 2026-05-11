//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what FetchBlock의 Params와 중첩 Fetch의 ParamBind를 재귀 수집한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func collectFetchParamBinds(f stmlparser.FetchBlock, params []stmlparser.ParamBind) []stmlparser.ParamBind {
	params = append(params, f.Params...)
	for _, child := range f.NestedFetches {
		params = collectFetchParamBinds(child, params)
	}
	return params
}
