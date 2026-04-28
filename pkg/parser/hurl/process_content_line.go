//ff:func feature=crosscheck type=util control=selection topic=scenario-check
//ff:what processContentLine — section 에 따라 content 라인을 수집기로 라우팅

package hurl

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
		collectCaptureLine(s, line)
	case "asserts":
		collectAssertLine(s, line)
	}
}
