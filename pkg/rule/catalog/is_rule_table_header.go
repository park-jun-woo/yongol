//ff:func feature=rule type=util control=sequence topic=catalog
//ff:what isRuleTableHeader — `| Rule ID | Level | Description | Source |` 헤더 검사
package catalog

import "strings"

// isRuleTableHeader checks whether the trimmed row is the canonical
// rule table header: "| Rule ID | Level | Description | Source |".
func isRuleTableHeader(row string) bool {
	cells := splitRow(row)
	if len(cells) < 4 {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(cells[0]), "Rule ID") &&
		strings.EqualFold(strings.TrimSpace(cells[1]), "Level") &&
		strings.EqualFold(strings.TrimSpace(cells[2]), "Description") &&
		strings.EqualFold(strings.TrimSpace(cells[3]), "Source")
}
