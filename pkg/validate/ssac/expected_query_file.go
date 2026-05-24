//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what expectedQueryFile — 모델명에서 기대 쿼리 파일 경로 역산 (RefreshToken → db/queries/refresh_tokens.sql)

package ssac

import (
	"github.com/jinzhu/inflection"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// expectedQueryFile derives the expected sqlc query file path from a model
// name. This is the reverse of modelFromFilename: PascalCase → snake_case →
// plural → append .sql. Example: "RefreshToken" → "db/queries/refresh_tokens.sql".
func expectedQueryFile(model string) string {
	snake := caseconv.PascalToSnake(model)
	plural := inflection.Plural(snake)
	return "db/queries/" + plural + ".sql"
}
