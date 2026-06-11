//ff:func feature=generate type=util control=sequence
//ff:what addDocumentTitle — 페이지 노드 하나의 document.title 기록 (라벨 폴백 = 페이지명, 앱명 결합, 최초 등장 우선)

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// addDocumentTitle records one sitemap node's document.title:
// "<label> · <app name>" (the label falls back to the page name for
// labelless nodes, the bare label when the manifest carries no app name).
// The first occurrence wins — a duplicate listing is TM-40's ERROR, so
// this is only a deterministic tie-break for the emit path. Nodes without
// a page (group labels, external links) record nothing.
func addDocumentTitle(node stmlparser.SitemapNode, appName string, titles map[string]string) {
	if node.Page == "" {
		return
	}
	if _, seen := titles[node.Page]; seen {
		return
	}
	label := node.Label
	if label == "" {
		label = node.Page
	}
	if appName == "" {
		titles[node.Page] = label
		return
	}
	titles[node.Page] = label + " · " + appName
}
