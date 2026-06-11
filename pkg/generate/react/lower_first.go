//ff:func feature=gen-react type=util control=sequence
//ff:what lowerFirst — 첫 글자만 소문자로 변환 (operationId → 데이터 변수명, 페이지 방출기의 toLowerFirst 규약)

package react

import "unicode"

// lowerFirst lowercases the first rune — the page emitter's data-variable
// naming convention (toLowerFirst in generate/react/stml): operationId
// "ListMyBuildings" names its useQuery data "listMyBuildingsData". The
// layout dynamic-group queries (plans/stml/sitemap Phase007) reuse it so
// the two emitters never disagree on the alias.
func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
