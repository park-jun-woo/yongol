//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what collectValidateInsertScan — validate 내부 INSERT 본문 수집 (unquoted `;` 종결)

package ddl

import (
	"strings"
)

// collectValidateInsertScan consumes lines starting at i (which matched
// insertIntoLineRe) until a top-level unquoted `;` terminates the
// statement. Returns the scan result plus the next index to resume from.
// The body is written verbatim including the terminator's line (the
// whole line, matching the pre-split behaviour).
func collectValidateInsertScan(lines []string, i int, table string, annotated bool) (insertScan, int) {
	var buf strings.Builder
	j := i
	done := false
	inSingle := false
	for j < len(lines) && !done {
		ln := lines[j]
		if j > i {
			buf.WriteByte('\n')
		}
		done, inSingle = scanValidateLineForTerminator(ln, inSingle)
		buf.WriteString(ln)
		j++
	}
	return insertScan{
		Table:     table,
		SQL:       buf.String(),
		StartLine: i + 1,
		Annotated: annotated,
	}, j
}
