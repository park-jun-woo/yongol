//ff:type feature=rule type=model
//ff:what StringSet — O(1) 조회를 위한 문자열 집합 별칭
package rule

// StringSet is a set of strings for O(1) lookup.
type StringSet = map[string]bool
