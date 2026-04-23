//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what scanHintComments — 단일 SQL 파일에서 `-- @<tag>` 힌트 주석 수집

package ddl

import (
	"bufio"
	"strings"
)

// scanHintComments reads a single SQL file line-by-line, maintaining CREATE
// TABLE context and pending standalone hints. Returns all extracted
// HintComment records; unknown tags are preserved (the caller decides).
func scanHintComments(r interface{ Read([]byte) (int, error) }, path string) ([]HintComment, error) {
	var out []HintComment
	sc := bufio.NewScanner(r.(interface{ Read([]byte) (int, error) }))
	lineNum := 0
	tableCtx := ""
	// pendingHints are hints on stand-alone comment lines that should
	// attach to the *next* non-blank DDL line.
	var pending []*HintComment
	for sc.Scan() {
		lineNum++
		ln := sc.Text()
		trim := strings.TrimSpace(ln)
		upper := strings.ToUpper(trim)
		if strings.HasPrefix(upper, "CREATE TABLE") {
			tableCtx, out, pending = handleCreateTableLine(trim, pending, out)
			continue
		}
		if strings.HasPrefix(trim, "--") {
			pending = appendPendingHint(trim, path, lineNum, tableCtx, pending)
			continue
		}
		ddlPart, comment := splitTrailingComment(ln)
		ddlTrim := strings.TrimSpace(ddlPart)
		if len(pending) > 0 && ddlTrim != "" {
			out, pending = drainPendingHints(pending, tableCtx, ddlTrim, out)
		}
		if comment != "" {
			out = applyTrailingCommentHint(comment, ddlTrim, path, lineNum, tableCtx, out)
		}
	}
	// Any pending hints without a following DDL line are dropped.
	return out, sc.Err()
}
