//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what modelForSQLCMethod — sqlc query 카탈로그에서 method name → 모델명 조회

package ssac

// modelForSQLCMethod returns QuerySpec.Model for the given sqlc query
// name. Empty when no match — caller falls back to direct emit.
func (g *methodGen) modelForSQLCMethod(method string) string {
	for _, q := range g.SQLcQueries {
		if q.Name == method {
			return q.Model
		}
	}
	return ""
}
