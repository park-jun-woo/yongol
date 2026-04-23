//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what parseTableItems — CREATE TABLE 본문을 top-level `,` 로 나눠 각 item 을 파싱
package migration

import "strings"

// parseTableItems splits the body of a CREATE TABLE on top-level commas
// and routes each item to parseTableItem.
func parseTableItems(t *Table, body string) {
	for _, it := range splitTopLevel(body, ',') {
		parseTableItem(t, strings.TrimSpace(it))
	}
}
