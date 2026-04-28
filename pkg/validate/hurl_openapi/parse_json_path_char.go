//ff:func feature=validate type=util control=selection topic=hurl-openapi
//ff:what parseJSONPathChar — parseJSONPath 루프의 1 글자 처리 (cur/out 갱신, 새로운 i 반환)

package hurl_openapi

import "strings"

// parseJSONPathChar processes a single byte inside parseJSONPath. It may
// flush the pending identifier into out, consume a `[index]` chunk, or
// append a regular byte to cur. Returns the new loop index.
func parseJSONPathChar(p string, i int, cur *strings.Builder, out *[]string) int {
	c := p[i]
	switch c {
	case '.':
		flushCurJSONPath(cur, out)
		return i
	case '[':
		flushCurJSONPath(cur, out)
		end := strings.Index(p[i:], "]")
		if end < 0 {
			// Return len(p)-1 so the enclosing for-loop terminates
			// (caller's `i++` moves us past the end-of-string guard).
			return len(p) - 1
		}
		*out = append(*out, p[i:i+end+1])
		return i + end
	default:
		cur.WriteByte(c)
		return i
	}
}
