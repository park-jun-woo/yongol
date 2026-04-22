//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what isSourceFile — 파일이 도구 원본 산출인지 (이미 분할된 결과 / 보존 파일 제외) 판정
package splitter

import (
	"bufio"
	"os"
	"strings"
)

// isSourceFile decides whether path should be fed through SplitFile. The
// check has two parts:
//
//  1. Name match — the basename must match the tool's source pattern
//     (matchesOriginal). Preserved files (querier.go / db.go) are
//     excluded upstream so this ignores them implicitly.
//  2. Provenance check — if the file's top lines contain a //ff:func or
//     //ff:type header, it is already a splitter output. Skip it.
//
// The provenance check makes SplitDirectory idempotent: stale split
// outputs from a previous run survive re-entry without being double
// split, and cleanOriginal will still remove them if they are obsolete
// (not in the new keep set).
func isSourceFile(path string, tool Tool) bool {
	base := tailSegment(path)
	if !matchesOriginal(base, tool) {
		return false
	}
	f, err := os.Open(path)
	if err != nil {
		return true
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	lines := 0
	for sc.Scan() && lines < 8 {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "//ff:func") || strings.HasPrefix(line, "//ff:type") {
			return false
		}
		lines++
	}
	return true
}
