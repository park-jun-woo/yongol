//ff:type feature=crosscheck type=parser topic=scenario-check
//ff:what parseState — .hurl 파싱 중 section 상태 추적 (request-headers / body / captures / asserts)

package hurl

import "strings"

// parseState tracks the line-by-line walk across hurl sections. Sections
// matter because the meaning of `jsonpath "$..."` differs between
// [Asserts] and [Captures].
type parseState struct {
	path    string
	lineNum int
	current *HurlEntry
	entries []HurlEntry
	section string // "", "request-headers", "body", "captures", "asserts"
	bodyBuf strings.Builder
}
