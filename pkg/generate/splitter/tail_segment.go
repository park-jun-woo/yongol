//ff:func feature=gen-splitter type=util control=sequence
//ff:what tailSegment — 파일 경로에서 basename(마지막 slash 뒤 구간)만 분리
package splitter

import "path/filepath"

// tailSegment returns the basename of a path. It is a thin wrapper kept
// in its own file so the call site in isSourceFile stays self-describing
// and the F1 one-func-per-file rule is honoured without introducing an
// unrelated import group in a neighbour.
func tailSegment(path string) string {
	return filepath.Base(path)
}
