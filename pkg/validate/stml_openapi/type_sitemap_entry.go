//ff:type feature=validate type=model topic=stml-openapi
//ff:what sitemapEntry — 사이트맵 트리를 평탄화한 노드 + 사람이 읽는 위치 경로 (TM-39/40/42 공유)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapEntry pairs one sitemap node with its human-readable position
// (e.g. `nav[0] > 건물 관리 > 건물 목록`) so the sitemap rules can point at
// the exact spot in the tree — sitemap.html carries no line information
// through the HTML parser.
type sitemapEntry struct {
	Node stml.SitemapNode
	Path string
}
