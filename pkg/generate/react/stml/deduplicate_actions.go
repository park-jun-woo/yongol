//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what OperationID 기준으로 중복 ActionBlock을 제거한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// deduplicateActions removes duplicate actions by OperationID.
func deduplicateActions(actions []stmlparser.ActionBlock) []stmlparser.ActionBlock {
	seen := map[string]bool{}
	var result []stmlparser.ActionBlock
	for _, a := range actions {
		if !seen[a.OperationID] {
			seen[a.OperationID] = true
			result = append(result, a)
		}
	}
	return result
}
