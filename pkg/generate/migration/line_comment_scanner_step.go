//ff:func feature=migration type=util control=selection
//ff:what lineCommentScanner.step — 한 바이트 처리, `--` 감지 시 hit=true
package migration

// step consumes line[i]. Returns (nextIdx, hit). When hit is true the
// caller should treat i as the start index of the `--` comment.
func (s *lineCommentScanner) step(line string, i int) (int, bool) {
	c := line[i]
	switch c {
	case '\'':
		return s.stepQuote(line, i), false
	case '-':
		if !s.inSQ && line[i+1] == '-' {
			return i, true
		}
	}
	return i, false
}
