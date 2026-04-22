//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what buildExpectedQueries — fs.SQLcQueries 에서 허용 query name/method 집합 구축

package contract

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildExpectedQueries returns the set of acceptable sqlc method names
// that a preserved handler may call. Both the raw `-- name:` ident
// and the prefix-stripped method form are accepted because different
// generators reach for different spellings.
func buildExpectedQueries(fs *yongol.Fullstack) map[string]bool {
	queries := map[string]bool{}
	for _, q := range fs.SQLcQueries {
		if q.Name != "" {
			queries[q.Name] = true
		}
		if q.Method != "" {
			queries[q.Method] = true
		}
	}
	return queries
}
