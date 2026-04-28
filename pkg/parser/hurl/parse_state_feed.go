//ff:func feature=crosscheck type=parser control=sequence topic=scenario-check
//ff:what parseState.feed — 한 줄을 소비하여 현재 entry 의 상태 전이

package hurl

import (
	"regexp"
	"strings"
)

var (
	reHurlRequest  = regexp.MustCompile(`^(GET|POST|PUT|DELETE|PATCH)\s+(?:\{\{host\}\}|https?://[^/]*)(/\S*)`)
	reHurlResponse = regexp.MustCompile(`^HTTP\s+(\d+)`)
)

func (s *parseState) feed(raw string) {
	line := strings.TrimSpace(raw)
	if m := reHurlRequest.FindStringSubmatch(line); m != nil {
		s.flushEntry()
		s.current = &HurlEntry{
			Method: m[1],
			Path:   trimQuery(m[2]),
			File:   s.path,
			Line:   s.lineNum,
		}
		s.section = "request-headers"
		s.bodyBuf.Reset()
		return
	}
	if s.current == nil {
		return
	}
	if m := reHurlResponse.FindStringSubmatch(line); m != nil {
		s.current.StatusCode = m[1]
		s.flushRequestBody()
		s.section = "response-headers"
		return
	}
	if processSectionHeader(s, line) {
		return
	}
	processContentLine(s, raw, line)
}
