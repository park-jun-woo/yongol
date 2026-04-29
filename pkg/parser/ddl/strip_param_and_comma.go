//ff:func feature=manifest type=util control=sequence
//ff:what stripParamAndComma — 토큰 끝의 "," 와 "(...)" 파라미터 리스트 제거

package ddl

import "strings"

// stripParamAndComma removes a trailing "(...)" parameter list and a
// trailing comma from a token. Used by matchMultiTokenHead when
// comparing the last token of a multi-word head against the matrix
// entry (e.g. "VARYING(255)," → "VARYING").
func stripParamAndComma(s string) string {
	s = strings.TrimSuffix(s, ",")
	if idx := strings.Index(s, "("); idx > 0 {
		s = s[:idx]
	}
	return s
}
