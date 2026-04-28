//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what handleRequestHeaderOrBodyStart — 요청 헤더 / JSON body 시작 구분

package hurl

import "strings"

// handleRequestHeaderOrBodyStart decides whether a line in the
// request-headers region is a header ("Name: value"), a comment / blank
// line, or the opening brace / bracket of a JSON body. In the last case
// the parser transitions to "body" and starts buffering.
func handleRequestHeaderOrBodyStart(s *parseState, raw, trimmed string) {
	if trimmed == "" || strings.HasPrefix(trimmed, "#") {
		return
	}
	if strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[") {
		s.section = "body"
		s.bodyBuf.WriteString(raw)
		s.bodyBuf.WriteByte('\n')
		return
	}
	m := reHeaderLine.FindStringSubmatch(trimmed)
	if m == nil || s.current == nil {
		return
	}
	s.current.Headers = append(s.current.Headers, HurlHeader{
		Name:  m[1],
		Value: m[2],
		Line:  s.lineNum,
	})
}
