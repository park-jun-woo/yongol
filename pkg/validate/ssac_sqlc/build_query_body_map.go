//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what buildQueryBodyMap — fs.SQLcQueries → query-name 키 / QuerySpec 값 맵 (XQS-20 lookup 용)

package ssac_sqlc

import (
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildQueryBodyMap returns a map from sqlc query name → its full QuerySpec.
// XQS-20 needs both Body (for RETURNING extraction) and Cardinality (skip
// `:exec`), so the whole spec is exposed rather than just one field.
func buildQueryBodyMap(fs *yongol.Fullstack) map[string]sqlcparser.QuerySpec {
	out := make(map[string]sqlcparser.QuerySpec, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		out[q.Name] = q
	}
	return out
}
