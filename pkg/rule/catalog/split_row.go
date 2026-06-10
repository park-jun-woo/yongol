//ff:func feature=rule type=util control=iteration dimension=1 topic=catalog
//ff:what splitRow — markdown 테이블 행을 비이스케이프 `|` 기준으로 분리, 셀 내 `\|` 는 `|` 로 복원
package catalog

import "strings"

// splitRow splits a markdown table row on unescaped pipe characters,
// stripping the leading and trailing pipes. Per GFM, a pipe inside a cell
// must be written as `\|`; such escaped pipes do not delimit cells and are
// restored to a literal `|` in the cell value (e.g. TM-17's `\|\|` code
// span, SEC-201's `cookie\|hybrid`). Any other backslash sequence is kept
// verbatim.
func splitRow(row string) []string {
	row = strings.TrimSpace(row)
	row = strings.TrimPrefix(row, "|")

	var cells []string
	var cell strings.Builder
	escaped := false
	for _, r := range row {
		switch {
		case escaped && r == '|':
			cell.WriteRune('|')
			escaped = false
		case escaped:
			cell.WriteByte('\\')
			cell.WriteRune(r)
			escaped = false
		case r == '\\':
			escaped = true
		case r == '|':
			cells = append(cells, cell.String())
			cell.Reset()
		default:
			cell.WriteRune(r)
		}
	}
	if escaped {
		cell.WriteByte('\\')
	}
	// A trailing unescaped pipe closes the row; it must not add an empty cell.
	if cell.Len() > 0 || !strings.HasSuffix(row, "|") {
		cells = append(cells, cell.String())
	}
	return cells
}
