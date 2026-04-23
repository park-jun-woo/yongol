//ff:func feature=util type=util control=iteration dimension=1 topic=string-convert
//ff:what SnakeToPascalSqlc — sqlc naming 규칙 ("id"/"ids" → 전체 대문자) snake→Pascal

package caseconv

import "strings"

// SnakeToPascalSqlc converts snake_case to PascalCase with sqlc's naming
// convention: "id"/"ids" parts become fully uppercase, other parts are
// capitalize-first. Example: "org_id" → "OrgID", "user_ids" → "UserIDS".
func SnakeToPascalSqlc(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		up := strings.ToUpper(p)
		if up == "ID" || up == "IDS" {
			b.WriteString(up)
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return b.String()
}
