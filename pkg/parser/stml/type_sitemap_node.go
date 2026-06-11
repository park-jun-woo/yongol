//ff:type feature=stml-parse type=model
//ff:what 사이트맵 트리의 <li> 노드 하나 (페이지/그룹 라벨/외부 링크 + 자식)

package stml

// SitemapNode is one <li> of the sitemap tree.
type SitemapNode struct {
	Page  string   // data-page ("" = group label or external link)
	Label string   // direct text of the <li>
	Href  string   // external link (<a href> child)
	Index bool     // data-index — "/" redirect target
	Menu  bool     // data-menu (default true; "false" hides the menu entry)
	Icon  string   // data-icon
	Roles []string // data-roles — menu entry visible only for these roles; the subtree inherits by nesting (Phase005)
	// CrumbField is data-crumb-field — the response field of the page's
	// first data-fetch whose value replaces the static crumb label (and
	// document.title) once the fetch arrives (Phase006). Page items only;
	// TM-39 rejects it on a group <li>, TM-50 validates the field.
	CrumbField string
	// Dynamic menu group (plans/stml/sitemap Phase007, DESIGN §4.11 (a)) —
	// the page vocabulary reused on the node's nested <ul>: the group's
	// items are the rows of an OpenAPI list response instead of static
	// children ("my buildings" workspace-switcher pattern). All five fields
	// come from the first nested <ul> carrying any of them.
	Fetch         string          // data-fetch — operationId whose response feeds the items
	Each          string          // data-each — response array field the items map over
	Link          string          // data-link — target page name of every item
	LinkParamsRaw string          // raw data-link-params value (TM-32 re-parses it for syntax diagnostics)
	LinkParams    []LinkParamBind // parsed bindings (empty when absent or syntactically invalid)
	LabelField    string          // data-label-field — item field rendered as the menu label (TM-30 validates)
	Children      []SitemapNode
}
