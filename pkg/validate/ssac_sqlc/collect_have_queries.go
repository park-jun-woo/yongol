//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what collectHaveQueries — fs.SQLcQueries 의 이름 집합 생성

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectHaveQueries returns the set of sqlc query names present in fs
// so XQS-19 can check built-in requirements in O(1).
func collectHaveQueries(fs *yongol.Fullstack) map[string]bool {
	have := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		have[q.Name] = true
	}
	return have
}
