//ff:func feature=gen-hurl type=util control=sequence
//ff:what writeStep — 단일 step을 hurl 텍스트 포맷으로 변환 (8단계 조립)
package hurl

import (
	"fmt"
	"strings"
)

// writeStep converts a single step into hurl text lines.
// 8-step assembly order: Comment → Request line → Query params →
// Authorization → Content-Type+body → HTTP status → [Captures] → [Asserts].
func writeStep(s step) string {
	var b strings.Builder
	if s.Comment != "" {
		b.WriteString("\n" + s.Comment + "\n")
	}
	b.WriteString(fmt.Sprintf("# %s\n", s.OperationID))

	reqLine := fmt.Sprintf("%s {{host}}%s", s.Method, s.Path)
	if s.QueryParams != "" {
		reqLine += "?" + s.QueryParams
	}
	b.WriteString(reqLine + "\n")

	if s.NeedsAuth {
		tok := s.TokenVar
		if tok == "" {
			tok = "token"
		}
		b.WriteString(fmt.Sprintf("Authorization: Bearer {{%s}}\n", tok))
	}
	if s.HasBody {
		b.WriteString("Content-Type: application/json\n")
		b.WriteString(s.BodyJSON + "\n")
	}
	b.WriteString(fmt.Sprintf("HTTP %d\n", s.StatusCode))
	b.WriteString(formatCaptures(s.Captures))
	b.WriteString(formatAssertions(s.Assertions))
	return b.String()
}
