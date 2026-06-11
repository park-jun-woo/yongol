//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what addSitemapCrumbFields — 사이트맵 노드 재귀 순회로 페이지별 data-crumb-field 누적 (지연 할당, 최초 등장 우선)

package stml

// addSitemapCrumbFields walks sitemap nodes depth-first, recording each
// page node's data-crumb-field (plans/stml/sitemap Phase006). The map is
// allocated lazily on the first declaration so a crumb-field-less sitemap
// yields nil — the byte-identity gate of every consumer. Page-less nodes
// (group labels, external links) record nothing even when they carry the
// attribute (TM-39 rejects that); the first occurrence of a page wins
// (a duplicate listing is TM-40's ERROR).
func addSitemapCrumbFields(nodes []SitemapNode, fields map[string]string) map[string]string {
	for _, n := range nodes {
		declares := n.Page != "" && n.CrumbField != ""
		if declares && fields == nil {
			fields = map[string]string{}
		}
		if _, seen := fields[n.Page]; declares && !seen {
			fields[n.Page] = n.CrumbField
		}
		fields = addSitemapCrumbFields(n.Children, fields)
	}
	return fields
}
