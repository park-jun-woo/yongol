//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sqlcPascalCase — snake_case → sqlc 컨벤션 PascalCase (ID/IDS 이니셜리즘 유지)

package ssac

import "strings"

// sqlcPascalCase converts snake_case to PascalCase following Go/sqlc acronym rules.
// "id" → "ID", "org_id" → "OrgID", "created_at" → "CreatedAt", "url" → "URL".
func sqlcPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		upper := strings.ToUpper(p)
		if sqlcAcronyms[upper] {
			b.WriteString(upper)
		} else {
			b.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return b.String()
}

// sqlcAcronyms: only acronyms that sqlc actually uppercases.
// sqlc treats "id" → "ID" but NOT "url" → "URL" (sqlc uses "Url").
var sqlcAcronyms = map[string]bool{
	"ID": true, "IDS": true,
}
