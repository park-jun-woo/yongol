//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what applySitemapDynamicGroup — 그룹 <li> 의 중첩 <ul> 에서 동적 그룹 어휘 5종을 SitemapNode 로 승격 (첫 동적 ul 승자)

package stml

import "golang.org/x/net/html"

// applySitemapDynamicGroup graduates the dynamic-group vocabulary of a
// sitemap group's nested <ul> onto its SitemapNode (plans/stml/sitemap
// Phase007, DESIGN §4.11 (a)): data-fetch / data-each / data-link /
// data-link-params / data-label-field. A <ul> without any of the five is
// a plain static container and contributes nothing; once a node has
// dynamic fields the first declaring <ul> won and later siblings are
// ignored. data-link-params keeps the raw value for TM-32's syntax
// re-parse and the parsed bindings when they parse cleanly (the LinkRef
// ParamsRaw/Params split).
func applySitemapDynamicGroup(ul *html.Node, node *SitemapNode) {
	hasAny := false
	for _, attr := range []string{"data-fetch", "data-each", "data-link", "data-link-params", "data-label-field"} {
		if hasAttr(ul, attr) {
			hasAny = true
			break
		}
	}
	if !hasAny {
		return
	}
	if node.Fetch != "" || node.Each != "" || node.Link != "" || node.LinkParamsRaw != "" || node.LabelField != "" {
		return
	}
	node.Fetch = getAttr(ul, "data-fetch")
	node.Each = getAttr(ul, "data-each")
	node.Link = getAttr(ul, "data-link")
	node.LinkParamsRaw = getAttr(ul, "data-link-params")
	node.LabelField = getAttr(ul, "data-label-field")
	if node.LinkParamsRaw == "" {
		return
	}
	if params, err := ParseLinkParams(node.LinkParamsRaw); err == nil {
		node.LinkParams = params
	}
}
