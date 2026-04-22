//ff:func feature=orchestrator type=util control=sequence dimension=1
//ff:what modelFromFilename — 파일명(users.sql)에서 단수화+PascalCase 모델명 도출
package sqlc

import (
	"strings"

	"github.com/jinzhu/inflection"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// modelFromFilename derives a PascalCase model name from a sqlc query
// filename. "users.sql" → "User", "user_profiles.sql" → "UserProfile".
func modelFromFilename(filename string) string {
	base := strings.TrimSuffix(filename, ".sql")
	return caseconv.SnakeToPascalSqlc(inflection.Singular(base))
}
