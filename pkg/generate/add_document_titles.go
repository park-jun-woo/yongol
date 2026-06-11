//ff:func feature=generate type=util control=iteration dimension=1
//ff:what addDocumentTitles — 사이트맵 노드 재귀 순회로 페이지별 타이틀 누적 (노드 단위 판정은 addDocumentTitle)

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// addDocumentTitles walks sitemap nodes depth-first, recording one
// document.title per page node through addDocumentTitle (label fallback,
// app-name join, first-occurrence-wins). Group labels and external links
// carry no page and add nothing of their own; their children still walk.
func addDocumentTitles(nodes []stmlparser.SitemapNode, appName string, titles map[string]string) {
	for _, n := range nodes {
		addDocumentTitle(n, appName, titles)
		addDocumentTitles(n.Children, appName, titles)
	}
}
