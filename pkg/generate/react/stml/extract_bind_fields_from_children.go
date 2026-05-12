//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what EachBlock의 Children에서 data-bind 필드명 목록을 추출한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// extractBindFieldsFromChildren extracts bind field names from ChildNode slice.
func extractBindFieldsFromChildren(nodes []stmlparser.ChildNode) []string {
	var fields []string
	for _, ch := range nodes {
		if ch.Kind == "bind" && ch.Bind != nil {
			fields = append(fields, ch.Bind.Name)
		}
	}
	return fields
}
