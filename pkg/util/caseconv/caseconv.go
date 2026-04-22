//ff:func feature=util type=util control=sequence topic=string-convert
//ff:what SnakeToPascal / SnakeToPascalSqlc / PascalToSnake / KebabToCamel — 공용 case 변환

package caseconv

import (
	"strings"

	"github.com/ettle/strcase"
)

// SnakeToPascal converts snake_case to PascalCase using the plain convention
// (capitalize-first per part). Example: "user_id" → "UserId", "per_page" → "PerPage".
func SnakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return b.String()
}

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

// PascalToSnake converts PascalCase / camelCase to snake_case via ettle/strcase.
func PascalToSnake(s string) string {
	return strcase.ToSnake(s)
}

// KebabToCamel converts kebab-case to camelCase. Strings without '-' are
// returned unchanged. Example: "data-fetch" → "dataFetch".
func KebabToCamel(s string) string {
	if !strings.Contains(s, "-") {
		return s
	}
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
