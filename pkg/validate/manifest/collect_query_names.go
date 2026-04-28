//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what collectQueryNames — fs.SQLcQueries 의 name 세트 구성

package manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectQueryNames returns the set of sqlc query names present in fs.
func collectQueryNames(fs *yongol.Fullstack) map[string]bool {
	out := make(map[string]bool, len(fs.SQLcQueries))
	for _, q := range fs.SQLcQueries {
		out[q.Name] = true
	}
	return out
}
