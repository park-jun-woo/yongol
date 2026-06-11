//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderSitemapDynamicGroup — 동적 메뉴 그룹 <li> 방출 (0개면 헤더째 비렌더, fetch 항목들을 NavLink 로 map, 정적 자식 선행)

package react

import (
	"fmt"
	"strings"
)

// renderSitemapDynamicGroup writes one dynamic menu group (plans/stml/
// sitemap Phase007): the whole <li> — header included — renders only while
// the fetched list has items, so loading, an empty response and an error
// all collapse to silent omission (a menu is not a content area; the page
// consuming the same operation owns the error's visibility). Items map
// over the useQuery data with the page each-emitter's conventions: `item`
// callback, the first item.* link-param source as the key (positional
// index without one), and a <NavLink end> whose to attribute came from the
// page data-link substitution (LinkToAttr) — active state is per item by
// route matching. Static children (if any) render before the dynamic items
// inside the same <ul>, through the ordinary item renderer.
func renderSitemapDynamicGroup(sb *strings.Builder, item sitemapMenuItem, indent string, rolesActive bool) {
	listExpr := fmt.Sprintf("(%sData?.%s ?? [])", lowerFirst(item.Fetch), strings.ReplaceAll(item.Each, ".", "?."))
	mapParams, keyExpr := "(item)", item.ItemKey
	if keyExpr == "" {
		mapParams, keyExpr = "(item, index)", "index"
	}
	fmt.Fprintf(sb, "%s{%s.length > 0 && (\n", indent, listExpr)
	fmt.Fprintf(sb, "%s  <li>\n", indent)
	fmt.Fprintf(sb, "%s    %s\n", indent, renderSitemapEntry(item))
	fmt.Fprintf(sb, "%s    <ul>\n", indent)
	for _, c := range item.Children {
		renderSitemapItem(sb, c, indent+"      ", rolesActive)
	}
	fmt.Fprintf(sb, "%s      {%s.map(%s => (\n", indent, listExpr, mapParams)
	fmt.Fprintf(sb, "%s        <li key={%s}><NavLink %s end>{item.%s}</NavLink></li>\n", indent, keyExpr, item.ItemToAttr, item.LabelField)
	fmt.Fprintf(sb, "%s      ))}\n", indent)
	fmt.Fprintf(sb, "%s    </ul>\n", indent)
	fmt.Fprintf(sb, "%s  </li>\n", indent)
	fmt.Fprintf(sb, "%s)}\n", indent)
}
