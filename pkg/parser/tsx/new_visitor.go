//ff:func feature=tsx-parser type=loader control=iteration dimension=1
//ff:what newVisitor — source bytes 에서 line offset 인덱스를 만들고 visitor 반환

package tsx

// newVisitor builds a line index from src so swc `span.start` byte offsets
// can be translated into (line, col). swc span starts at 1, not 0, per the
// swc documentation — adjusted here to standard 0-based offsets first.
func newVisitor(src []byte, page *PageSpec) *visitor {
	// lineOffset[0] = 0 (line 1 starts at byte 0)
	off := []int{0}
	for i, b := range src {
		if b == '\n' {
			off = append(off, i+1)
		}
	}
	return &visitor{src: src, lineOffset: off, page: page}
}
