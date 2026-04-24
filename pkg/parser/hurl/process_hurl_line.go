//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what 단일 Hurl 파일 라인을 파싱하여 요청/응답 상태 갱신 (section 분기 포함)
package hurl

import (
	"regexp"
	"strings"
)

var (
	reSectionHdr   = regexp.MustCompile(`^\[([A-Za-z][A-Za-z0-9]*)\]\s*$`)
	reHeaderLine   = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9\-]*):\s*(.*)$`)
	reJSONPath     = regexp.MustCompile(`jsonpath\s+"([^"]+)"`)
	reCaptureLine  = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s*:\s*(.+)$`)
	reHeaderSource = regexp.MustCompile(`^header\s+"([^"]+)"`)
)

// processSectionHeader detects `[Captures]` / `[Asserts]` / other
// section markers and switches the parser state accordingly. Returns
// true when the line was consumed by the section-header branch.
func processSectionHeader(s *parseState, line string) bool {
	m := reSectionHdr.FindStringSubmatch(line)
	if m == nil {
		return false
	}
	switch strings.ToLower(m[1]) {
	case "captures":
		s.section = "captures"
	case "asserts":
		s.section = "asserts"
	case "options", "querystringparams", "formparams", "multipartformdata", "cookies":
		s.section = "other"
	default:
		s.section = "other"
	}
	return true
}

// processContentLine routes a non-section, non-request, non-status line
// to the correct collector for the current parser state. Comment lines
// and blank lines are skipped where they would otherwise contribute
// noise (captures / asserts / headers); the body buffer preserves blank
// lines because hurl ends the body by seeing `HTTP <status>` rather
// than relying on blanks.
func processContentLine(s *parseState, raw, line string) {
	switch s.section {
	case "request-headers":
		handleRequestHeaderOrBodyStart(s, raw, line)
	case "body":
		s.bodyBuf.WriteString(raw)
		s.bodyBuf.WriteByte('\n')
	case "captures":
		if line == "" || strings.HasPrefix(line, "#") {
			return
		}
		c, ok := parseCaptureLine(line, s.lineNum)
		if ok && s.current != nil {
			s.current.Captures = append(s.current.Captures, c)
		}
	case "asserts":
		if line == "" || strings.HasPrefix(line, "#") {
			return
		}
		if jp, ok := parseAssertLine(line, s.lineNum); ok && s.current != nil {
			s.current.Asserts = append(s.current.Asserts, jp)
		}
	}
}

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
	if m := reHeaderLine.FindStringSubmatch(trimmed); m != nil && s.current != nil {
		s.current.Headers = append(s.current.Headers, HurlHeader{
			Name:  m[1],
			Value: m[2],
			Line:  s.lineNum,
		})
	}
}

// parseAssertLine extracts the first jsonpath argument in a line like
// `jsonpath "$.user.id" isInteger`. Returns false when the line lacks a
// jsonpath expression.
func parseAssertLine(line string, lineNum int) (HurlAssert, bool) {
	m := reJSONPath.FindStringSubmatch(line)
	if m == nil {
		return HurlAssert{}, false
	}
	return HurlAssert{JSONPath: m[1], Line: lineNum}, true
}

// parseCaptureLine parses a single [Captures] entry. Two forms are
// recognised: `<var>: jsonpath "$.x"` and `<var>: header "X-Y"`; any
// other source expression is stored verbatim so rule authors can
// special-case later.
func parseCaptureLine(line string, lineNum int) (HurlCapture, bool) {
	m := reCaptureLine.FindStringSubmatch(line)
	if m == nil {
		return HurlCapture{}, false
	}
	c := HurlCapture{Var: m[1], Line: lineNum}
	expr := strings.TrimSpace(m[2])
	if jp := reJSONPath.FindStringSubmatch(expr); jp != nil {
		c.Source = "jsonpath"
		c.JSONPath = jp[1]
		return c, true
	}
	if h := reHeaderSource.FindStringSubmatch(expr); h != nil {
		c.Source = "header"
		c.Header = h[1]
		return c, true
	}
	c.Source = expr
	return c, true
}
