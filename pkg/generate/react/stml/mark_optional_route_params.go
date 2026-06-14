//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지의 action-only route.* ParamBind에 Optional=true를 표시한다 (fetch 소비 param은 required로 남김, BUG-136)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// markOptionalRouteParams flags every action-only route.<Name> ParamBind in
// the page with Optional=true, mirroring collectRouteParams's required/optional
// rule (route params consumed by a data-fetch are required ":Name"; params
// consumed only by data-action blocks are optional ":Name?"). The flag rides on
// the ParamBind so renderParamArgs / renderRowMutateArg / renderUseQuery can
// emit a null guard for optional integer params without threading the optional
// set through their signatures (BUG-136). Must run before populateRowActionArgs
// so row-action mutate args capture the flag.
func markOptionalRouteParams(page *stmlparser.PageSpec) {
	required := requiredRouteNames(*page)
	for i := range page.Actions {
		setBindsOptional(page.Actions[i].Params, required)
	}
	markChildActionBindsOptional(page.Children, required)
}
