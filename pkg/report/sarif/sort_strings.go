//ff:func feature=report type=util control=iteration dimension=2 topic=sarif
//ff:what sortStrings — fallback rules 정렬용 insertion sort (sort 패키지 의존 회피)
package sarif

// sortStrings sorts s in place. Tiny helper to avoid a package-level
// `import "sort"` when buildDriverRules is exercised only in fallback mode.
func sortStrings(s []string) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
