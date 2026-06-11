//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what appendSitemapEntries — 사이트맵 노드 목록을 재귀 순회하며 위치 경로를 누적해 entries 에 추가

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendSitemapEntries walks nodes depth-first, extending the position path
// with each node's label (falling back to its page name, then "(group)").
func appendSitemapEntries(nodes []stml.SitemapNode, prefix string, entries *[]sitemapEntry) {
	for _, n := range nodes {
		display := n.Label
		if display == "" {
			display = n.Page
		}
		if display == "" {
			display = "(group)"
		}
		path := prefix + " > " + display
		*entries = append(*entries, sitemapEntry{Node: n, Path: path})
		appendSitemapEntries(n.Children, path, entries)
	}
}
