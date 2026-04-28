//ff:func feature=crosscheck type=util control=sequence topic=scenario-check
//ff:what collectCaptureLine — [Captures] 한 줄을 파싱해 current entry 에 append

package hurl

import "strings"

// collectCaptureLine parses one [Captures] line and appends the result
// into the current entry. Comment / blank lines are skipped.
func collectCaptureLine(s *parseState, line string) {
	if line == "" || strings.HasPrefix(line, "#") {
		return
	}
	c, ok := parseCaptureLine(line, s.lineNum)
	if !ok || s.current == nil {
		return
	}
	s.current.Captures = append(s.current.Captures, c)
}
