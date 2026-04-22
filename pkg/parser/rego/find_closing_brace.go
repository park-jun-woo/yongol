//ff:func feature=policy type=util control=iteration dimension=1
//ff:what findClosingBrace — 중괄호 깊이 추적하여 닫는 중괄호 인덱스 반환
package rego

func findClosingBrace(s string) int {
	depth := 1
	for i, c := range s {
		if c == '{' {
			depth++
		}
		if c == '}' {
			depth--
		}
		if depth == 0 {
			return i
		}
	}
	return -1
}
