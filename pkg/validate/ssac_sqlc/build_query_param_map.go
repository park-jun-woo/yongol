//ff:func feature=validate type=util control=iteration dimension=1 topic=sqlc
//ff:what buildQueryParamMap — sqlc QuerySpec에서 쿼리명 → Params set 맵 구축

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildQueryParamMap returns a map from sqlc query name → set of Params field names.
// e.g. "AuditLogListByOrgIDPaged" → {"OrgID": true, "Page": true, "PerPage": true, ...}
func buildQueryParamMap(fs *yongol.Fullstack) map[string]map[string]bool {
	result := make(map[string]map[string]bool)
	for _, q := range fs.SQLcQueries {
		if len(q.Params) == 0 {
			continue
		}
		set := make(map[string]bool, len(q.Params))
		for _, p := range q.Params {
			set[p] = true
		}
		result[q.Name] = set
	}
	return result
}
