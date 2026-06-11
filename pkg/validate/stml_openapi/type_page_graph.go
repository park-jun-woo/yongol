//ff:type feature=validate type=model topic=stml-openapi
//ff:what pageGraph — 도달성 검증용 페이지 그래프 (노드=전체 페이지, 간선=메뉴 렌더/링크/리다이렉트/브레드크럼 상행, 루트)

package stml_openapi

// pageGraph is the reachability graph of plans/stml/sitemap Phase002
// (DESIGN §4.10 — listing is a node, not an edge). Nodes are every STML
// page regardless of sitemap listing; edges exist only where the user can
// actually move: menu-rendered sitemap entries (folded into Roots — the
// menu hangs off the virtual root and is visible from everywhere),
// data-link declarations, resolvable data-redirect targets and the
// Phase004 breadcrumb up-links (edge (d) — child page → MenuRenderable
// sitemap ancestor). InSitemap and MenuBlocked carry the cause
// classification TM-43 puts in its message (listed-but-not-menu-rendered
// vs not listed at all).
type pageGraph struct {
	Pages       []string            // every STML page name, document order
	Roots       map[string]bool     // BFS seed: data-index ∪ data-entry pages ∪ manifest.frontend.index ∪ "/" mounts ∪ menu-rendered sitemap pages
	Edges       map[string][]string // source page → target pages (data-link / resolvable data-redirect / breadcrumb up-link)
	InSitemap   map[string]bool     // page names listed anywhere in the sitemap
	MenuBlocked map[string]string   // listed but non-menu-rendered page → human-readable reason
}
