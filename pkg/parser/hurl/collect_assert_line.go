//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what collectAssertLine — [Asserts] 한 줄을 파싱해 current entry 에 append

package hurl

import "strings"

// collectAssertLine parses one [Asserts] jsonpath line and appends the
// result into the current entry. Comment / blank lines are skipped.
func collectAssertLine(s *parseState, line string) {
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}
	jp, ok := parseAssertLine(line, s.lineNum)
	if !ok || s.current == nil {
		return
	}
	s.current.Asserts = append(s.current.Asserts, jp)
}
