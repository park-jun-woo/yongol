//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what PascalCase 문자열을 공백 구분 라벨로 변환한다
package stml

import "unicode"

func pascalToLabel(name string) string {
	if name == "" {
		return ""
	}
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	var words []string
	start := 0
	for i := 1; i < len(runes); i++ {
		if !isWordBoundary(runes, i) {
			continue
		}
		words = append(words, string(runes[start:i]))
		start = i
	}
	words = append(words, string(runes[start:]))
	return joinWords(words)
}
