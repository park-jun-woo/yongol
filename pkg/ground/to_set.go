//ff:func feature=rule type=util control=iteration dimension=1
//ff:what toSet — 문자열 슬라이스를 StringSet으로 변환
package ground

import "github.com/park-jun-woo/yongol/pkg/rule"

func toSet(vals []string) rule.StringSet {
	s := make(rule.StringSet, len(vals))
	for _, v := range vals {
		s[v] = true
	}
	return s
}
