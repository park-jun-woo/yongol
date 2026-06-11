//ff:type feature=gen-react type=model
//ff:what sitemapMenuItem — 사이트맵에서 파생된 레이아웃 메뉴 항목 하나 (페이지/그룹 라벨/외부 링크 + 자식)

package react

// sitemapMenuItem is one rendered entry of the sitemap-derived layout menu
// (plans/stml/sitemap Phase003). Only nodes the shared
// stml_openapi.MenuRenderable judgment admits become items — validation
// and emission never disagree on what renders.
type sitemapMenuItem struct {
	Kind     string   // "page" | "group" | "external"
	Label    string   // direct text of the sitemap <li>
	Icon     string   // lucide-react component name (PascalCase), "" = no icon
	To       string   // resolved route path (Kind "page")
	Href     string   // external URL (Kind "external")
	Prefixes []string // static route prefixes of menu-hidden descendants — ancestor active highlight
	Roles    []string // data-roles allowlist — item (subtree by nesting) renders only for these roles (Phase005)
	// Dynamic menu group (plans/stml/sitemap Phase007) — set only when the
	// sitemap node passes stml_openapi.DynamicMenuGroup (the complete
	// vocabulary set): the rendered items are the rows of the layout
	// useQuery for Fetch, mapped over Each, labeled by LabelField, each a
	// NavLink whose to attribute is ItemToAttr. Fetch == "" keeps the item
	// fully static.
	Fetch      string // data-fetch operationId — the layout useQuery key/call
	Each       string // data-each response array field the items map over
	LabelField string // item field rendered as the NavLink label
	ItemToAttr string // NavLink to attribute (LinkToAttr — `to="…"` or `to={…}`)
	ItemKey    string // map key expression ("item.<Field>"); "" = positional index
	Children   []sitemapMenuItem
}
