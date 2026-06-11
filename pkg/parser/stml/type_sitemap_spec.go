//ff:type feature=stml-parse type=model
//ff:what sitemap.html 파싱 결과 루트 — nav 블록 목록

package stml

// SitemapSpec represents the parsed frontend/sitemap.html — the central
// site-structure declaration (plans/stml/sitemap Phase001). Absent file =
// nil *SitemapSpec on the Fullstack; every behavior keyed off it stays off.
// The Phase001 reserved-attribute record (TM-45) is gone: every declared
// attribute graduated to a first-class field by Phase007.
type SitemapSpec struct {
	FileName string       // "sitemap.html"
	Navs     []SitemapNav // one per <nav data-sitemap> block, document order
}
