//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectPageEdges — 한 페이지의 나가는 간선 수집 (data-link 대상 + 해석 가능한 data-redirect 대상)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageEdges gathers the outgoing reachability edges of one page —
// DESIGN §4.10 edges (b) and (c): every data-link target on the page
// (collectLinkRefs, the TM-31 walk — row links included) and every
// resolvable data-redirect target of its actions (page.Actions plus the
// each-row actions the parser keeps on their blocks). Targets are raw
// page names; buildPageGraph drops the nonexistent ones (TM-26/31's
// findings).
func collectPageEdges(page stml.PageSpec, pages []stml.PageSpec, indexPages []string) []string {
	var targets []string
	var refs []linkRefCtx
	collectLinkRefs(page.Children, "", nil, false, nil, &refs)
	for _, ref := range refs {
		targets = append(targets, ref.Link.TargetPage)
	}
	actions := append([]stml.ActionBlock{}, page.Actions...)
	collectEachRowActions(page.Children, &actions)
	for _, a := range actions {
		targets = append(targets, resolveRedirectTargets(a.Redirect, pages, indexPages)...)
	}
	return targets
}
