//ff:func feature=rule type=util control=sequence
//ff:what mergedSet — Ground.Lookup[key] 의 기존 StringSet 을 반환하거나 없으면 새로 생성 (도메인 루프 누적용)
package ground

import "github.com/park-jun-woo/yongol/pkg/rule"

// mergedSet returns the existing StringSet stored at key so callers can union
// into it across a per-domain loop; when absent (or nil) a fresh set is
// returned. The caller is responsible for assigning the set back under key.
func mergedSet(g *rule.Ground, key string) rule.StringSet {
	if s, ok := g.Lookup[key]; ok && s != nil {
		return s
	}
	return make(rule.StringSet)
}
