//ff:type feature=stml-parse type=model
//ff:what 사이트맵의 <nav data-sitemap> 블록 하나 (layout/entry + 페이지 트리)

package stml

// SitemapNav is one <nav data-sitemap> block of the sitemap file.
type SitemapNav struct {
	Layout string // data-layout ("" = defaultLayout delegation)
	Entry  bool   // data-entry — every page in the block is a reachability root
	Items  []SitemapNode
}
