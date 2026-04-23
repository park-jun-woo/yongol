//ff:func feature=migration type=util control=selection
//ff:what lineCommentScanner.stepQuote — `'` 입력 시 escape ('') 또는 flag toggle 처리
package migration

// stepQuote handles a single-quote char at line[i].
func (s *lineCommentScanner) stepQuote(line string, i int) int {
	if s.inSQ && i+1 < len(line) && line[i+1] == '\'' {
		return i + 1
	}
	s.inSQ = !s.inSQ
	return i
}
