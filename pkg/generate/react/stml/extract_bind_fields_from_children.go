//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what EachBlock의 Children에서 data-bind 항목(태그 포함)을 추출한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// extractBindFieldsFromChildren extracts bind fields (with their tags) from a
// ChildNode slice. The full FieldBind is kept — not just the name — so a
// data-each cell can honor an <img data-bind> as a media cell instead of
// demoting it to text (plans/gen/frontend Phase037, BUG-126).
func extractBindFieldsFromChildren(nodes []stmlparser.ChildNode) []stmlparser.FieldBind {
	var fields []stmlparser.FieldBind
	for _, ch := range nodes {
		if ch.Kind == "bind" && ch.Bind != nil {
			fields = append(fields, *ch.Bind)
		}
	}
	return fields
}
