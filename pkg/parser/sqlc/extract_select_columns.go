//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what extractSelectColumns — SQL 본문에서 SELECT 절의 컬럼 목록 또는 SELECT * 여부를 추출

package sqlc

import (
	"regexp"
	"strings"
)

// selectClauseRe extracts the column part between SELECT and FROM.
// Case-insensitive, DOTALL so multi-line bodies match.
var selectClauseRe = regexp.MustCompile(`(?is)\bSELECT\s+(.*?)\bFROM\b`)

// extractSelectColumns analyses a SQL body to determine whether the query
// uses SELECT * or an explicit column list. For non-SELECT statements
// (INSERT, UPDATE, DELETE without a leading SELECT) both selectStar and
// selectCols are left at their zero values.
//
// When the SELECT clause is `*`, selectStar is true and selectCols is nil.
// When the clause lists specific columns, selectStar is false and
// selectCols contains the snake_case column names (aliases resolved to the
// alias name, table-qualified names stripped to the bare column, expressions
// ignored).
func extractSelectColumns(body string) (selectStar bool, selectCols []string) {
	m := selectClauseRe.FindStringSubmatch(body)
	if m == nil {
		return false, nil
	}
	colPart := strings.TrimSpace(m[1])

	// SELECT * — including table-qualified like t.*
	if colPart == "*" || strings.HasSuffix(colPart, ".*") {
		return true, nil
	}

	// Split by comma, parse each column expression.
	rawCols := strings.Split(colPart, ",")
	for _, raw := range rawCols {
		col := parseOneSelectColumn(strings.TrimSpace(raw))
		if col != "" {
			selectCols = append(selectCols, col)
		}
	}
	if len(selectCols) == 0 {
		// Could not extract any column names — treat as unknown.
		return false, nil
	}
	return false, selectCols
}

