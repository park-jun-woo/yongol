//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what collectSentinelInsert — INSERT 시작부터 종결 `;` 까지 본문을 수집 (따옴표 인식)
package ddl

import (
	"strings"
)

// collectSentinelInsert walks lines starting at index i (which matched
// insertIntoRe), consumes every subsequent line until a top-level `;` is
// found (single-quoted literals are respected) and returns the assembled
// scan result plus the next index to resume iteration at. Extracted
// from parseSentinelInserts so the outer walker stays at dimension 2.
func collectSentinelInsert(lines []string, i int, table string, annotated bool) (sentinelScanResult, int) {
	var buf strings.Builder
	j := i
	done := false
	inSingle := false
	for j < len(lines) && !done {
		ln := lines[j]
		if j > i {
			buf.WriteByte('\n')
		}
		buf.WriteString(ln)
		k, terminated := findUnquotedSemicolon(ln, inSingle)
		if terminated {
			rebuildBufferWithTerminator(&buf, lines, i, j, ln, k)
			done = true
		} else {
			inSingle = trackQuoteState(ln, inSingle)
		}
		j++
	}
	return sentinelScanResult{
		Table:     table,
		SQL:       buf.String(),
		StartLine: i + 1,
		Annotated: annotated,
	}, j
}
