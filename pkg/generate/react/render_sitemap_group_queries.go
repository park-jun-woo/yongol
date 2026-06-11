//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderSitemapGroupQueries — 동적 그룹 op 들의 useQuery const 들을 순서대로 방출 (renderLayoutTSX 의 루프 본체)

package react

import "strings"

// renderSitemapGroupQueries writes one layout useQuery const per dynamic
// menu group operation, in document order (plans/stml/sitemap Phase007) —
// renderLayoutTSX's loop body, one renderSitemapGroupQuery per op.
func renderSitemapGroupQueries(sb *strings.Builder, ops []string, bearerGate bool) {
	for _, op := range ops {
		renderSitemapGroupQuery(sb, op, bearerGate)
	}
}
