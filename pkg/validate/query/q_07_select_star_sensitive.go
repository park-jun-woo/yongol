//ff:func feature=validate type=rule control=iteration dimension=2 topic=query-structural
//ff:what Q-07 — WARNING when SELECT * is used on a table that has @sensitive columns

package query

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// selectStarFromRe matches `SELECT ... * ... FROM <table>` roughly.
var selectStarFromRe = regexp.MustCompile(`(?i)SELECT\s+\*\s+FROM\s+([A-Za-z_][A-Za-z0-9_]*)`)

// q07SelectStarSensitive validates Q-07: a SELECT * from a table that has any
// @sensitive column emits a WARNING. The star expands to the sensitive column
// and may leak it into the response. Use explicit column lists or re-check
// the @sensitive marker. Escape with `-- @allow-sensitive` when intentional.
func q07SelectStarSensitive(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.DDLTables) == 0 {
		return nil
	}
	sensitiveByTable := buildSensitiveColumnIndex(fs.DDLTables)
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		body, err := readQueryBody(q)
		if err != nil || body == nil {
			continue
		}
		if body.Escapes["@allow-sensitive"] {
			continue
		}
		matches := selectStarFromRe.FindAllStringSubmatch(body.Text, -1)
		for _, m := range matches {
			table := strings.ToLower(m[1])
			cols := sensitiveByTable[table]
			if len(cols) == 0 {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    q.File,
				Line:    q.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[Q-07] SELECT * in %s exposes @sensitive columns of %s: %s", q.Name, table, strings.Join(cols, ", ")),
				Advice:  "Replace with an explicit column list, or add `-- +allow-sensitive` if intentional exposure is accepted",
			})
		}
	}
	return diags
}
