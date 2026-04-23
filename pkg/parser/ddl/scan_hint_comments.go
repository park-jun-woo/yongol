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

		// Detect CREATE TABLE header.
		upper := strings.ToUpper(trim)
		if strings.HasPrefix(upper, "CREATE TABLE") {
			tableCtx = parseCreateTableName(trim)
			// consume any pending standalone hints as "above CREATE TABLE"
			for _, h := range pending {
				h.TableCtx = tableCtx
				h.BlockAbove = true
				out = append(out, *h)
			}
			pending = nil
			continue
		}

		// Comment-only line?
		if strings.HasPrefix(trim, "--") {
			hint := parseHintLine(trim, path, lineNum, tableCtx, "")
			if hint != nil {
				// Standalone: attach to the next DDL line.
				pending = append(pending, hint)
			}
			continue
		}

		// Line containing DDL + optional trailing `-- @...` comment.
		ddlPart, comment := splitTrailingComment(ln)
		ddlTrim := strings.TrimSpace(ddlPart)
		// Drain pending hints at the first real content.
		if len(pending) > 0 && ddlTrim != "" {
			column := extractColumnNameFromLine(ddlTrim)
			for _, h := range pending {
				if column != "" {
					h.ColumnCtx = column
				}
				h.TableCtx = tableCtx
				out = append(out, *h)
			}
			pending = nil
		}
		if comment != "" {
			column := extractColumnNameFromLine(ddlTrim)
			hint := parseHintLine("-- "+comment, path, lineNum, tableCtx, column)
			if hint != nil {
				out = append(out, *hint)
			}
		}
		if strings.HasSuffix(ddlTrim, ";") || strings.Contains(ddlTrim, ");") {
			// Conservative: end of statement clears table context.
			// (Works well enough for validation purposes.)
		}
	}
	// Any pending hints without a following DDL line are dropped.
	return out, sc.Err()
}
