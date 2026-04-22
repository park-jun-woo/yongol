//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sqlcFieldName — DDL 컬럼명을 sqlc 생성 Go 필드 이름으로 매핑 (ID 이니셜리즘 반영)

package sqlcpost

import "strings"

// sqlcFieldName maps a DDL column name to the sqlc-generated Go field name.
// Matches the observable convention in generated models: snake_case →
// PascalCase (e.g. org_id → OrgID, password_hash → PasswordHash, url → Url).
// sqlc's initialism handling treats "ID" specially.
func sqlcFieldName(col string) string {
	parts := strings.Split(col, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		up := strings.ToUpper(p)
		// sqlc initialism: "id" alone at any position becomes "ID".
		if up == "ID" {
			b.WriteString("ID")
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}
