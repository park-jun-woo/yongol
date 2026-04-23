//ff:func feature=gen-gogin type=util control=sequence topic=path-convert
//ff:what openAPIPathStep — openAPIPathToGin 한 스텝: 현재 위치에서 '{name}' 또는 단일 바이트 처리

package boot

import "strings"

// openAPIPathStep advances one step through the OpenAPI path at position i,
// writing either the translated gin segment (`:name`) or a single verbatim
// byte to b. Returns the next cursor position so the parent loop stays at
// depth 1.
func openAPIPathStep(p string, i int, b *strings.Builder) int {
	if p[i] != '{' {
		b.WriteByte(p[i])
		return i + 1
	}
	end := strings.IndexByte(p[i:], '}')
	if end <= 0 {
		b.WriteByte(p[i])
		return i + 1
	}
	b.WriteByte(':')
	b.WriteString(p[i+1 : i+end])
	return i + end + 1
}
